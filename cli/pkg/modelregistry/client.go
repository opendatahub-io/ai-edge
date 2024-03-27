/*
Copyright 2024. Open Data Hub Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package modelregistry

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/kubeflow/model-registry/pkg/openapi"
)

type StringList struct {
	Items []string `json:"items"`
}

// Client is a client for the model registry service
type Client struct {
	// modelRegistryURL is the URL of the model registry service
	modelRegistryURL    string
	modelRegistryClient *openapi.APIClient
}

// NewClient creates a new Client
func NewClient(modelRegistryURL string) *Client {
	configuration := openapi.NewConfiguration()
	configuration.Servers = openapi.ServerConfigurations{
		{
			URL: modelRegistryURL,
		},
	}

	return &Client{
		modelRegistryURL:    modelRegistryURL,
		modelRegistryClient: openapi.NewAPIClient(configuration),
	}
}

// AutoRegisterModelVersionArtifact is a convenience method to create a registered model, model version and model
// artifact in one call
//
// Errors:
// - ErrModelExists is returned when the model already exists
// - ErrVersionExists is returned when the version already exists
// - ErrArtifactExists is returned when the image already exists
func (c *Client) AutoRegisterModelVersionArtifact(
	modelName, modelDescription, versionName, artifactName, artifactExternalId, uri, modelFormatName, modelFormatVersion string,
	metaData map[string]openapi.MetadataValue,
) (
	*openapi.RegisteredModel,
	*openapi.ModelVersion,
	*openapi.Artifact,
	error,
) {
	if modelName == "" || modelDescription == "" || versionName == "" {
		return nil, nil, nil, fmt.Errorf("name, description and version are required")
	}

	m, err := c.CreateRegisteredModel(modelName, modelDescription, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	v, err := c.CreateModelVersion(m.GetId(), versionName, metaData)
	if err != nil {
		return nil, nil, nil, err
	}

	a, err := c.CreateModelArtifact(
		v.GetId(), artifactName, modelDescription, uri, modelFormatName, modelFormatVersion, artifactExternalId,
	)

	if err != nil {
		return nil, nil, nil, err
	}

	return m, v, a, nil
}

func (c *Client) CreateRegisteredModel(name string, description string, metaData map[string]openapi.MetadataValue) (
	*openapi.RegisteredModel,
	error,
) {
	if name == "" || description == "" {
		return nil, fmt.Errorf("name and description are required")
	}

	m := openapi.NewRegisteredModelCreateWithDefaults()
	m.SetName(name)
	m.SetDescription(description)

	if metaData != nil {
		m.SetCustomProperties(metaData)
	}

	model, resp, err := c.modelRegistryClient.ModelRegistryServiceAPI.CreateRegisteredModel(context.Background()).
		RegisteredModelCreate(*m).Execute()
	if err != nil {
		if resp == nil {
			return nil, fmt.Errorf("error creating registered model: %w", err)
		} else {
			// Currently model registry returns 500 when the model exists. This is a workaround to handle the error
			// until the model registry is fixed. The workaround is to check the error message and return ErrModelExists
			// if the error message contains the expected error message.
			// TODO: Remove this workaround when model registry returns 403 when the model exists
			if resp.StatusCode != 201 && isOpenAPIErrorOfKind(err, ErrAlreadyExists) {
				return nil, fmt.Errorf("%w. model name: %s", ErrModelExists, name)
			}
			// This is a weird case where we got a response and an error that we're unable to handle.
			return nil, fmt.Errorf(
				"error while creating a registered model: server responded with %s %w", resp.Status, err,
			)
		}
	}
	if resp.StatusCode != 201 {
		return nil, fmt.Errorf("failed to create a registered model: %s", resp.Status)
	}
	return model, nil
}

func (c *Client) CreateModelVersion(
	modelId string,
	versionName string,
	metaData map[string]openapi.MetadataValue,
) (*openapi.ModelVersion, error) {

	if modelId == "" || versionName == "" {
		return nil, fmt.Errorf("model ID and version are required")
	}

	modelVersion := openapi.NewModelVersionWithDefaults()
	modelVersion.SetName(versionName)

	if metaData != nil {
		modelVersion.SetCustomProperties(metaData)
	}

	v, resp, err := c.modelRegistryClient.ModelRegistryServiceAPI.CreateRegisteredModelVersion(
		context.Background(), modelId,
	).ModelVersion(*modelVersion).Execute()

	if err != nil {
		if resp == nil {
			return nil, fmt.Errorf("error creating model version: %w", err)
		} else {
			if resp.StatusCode != 201 {
				// TODO: Remove this workaround when model registry returns 404 when the model is not found
				if isOpenAPIErrorOfKind(err, ErrModelNotFound) {
					return nil, fmt.Errorf("%w. model id: %s", ErrModelNotFound, modelId)
				}
				if isOpenAPIErrorOfKind(err, ErrAlreadyExists) {
					return nil, fmt.Errorf("%w. model id: %s version name: %s", ErrVersionExists, modelId, versionName)
				}
			}
			// This is a weird case where we got a response and an error that we're unable to handle.
			return nil, fmt.Errorf(
				"error while creating a registered model version: server responded with %s %w", resp.Status, err,
			)
		}
	}

	if resp.StatusCode != 201 {
		return nil, fmt.Errorf("failed to create a registered model version: %s", resp.Status)
	}
	return v, nil
}

func (c *Client) CreateModelArtifact(
	versionId string,
	artifactName string,
	description string,
	uri string,
	modelFormatName string,
	modelFormatVersion string,
	externalId string,
) (*openapi.Artifact, error) {
	if versionId == "" || artifactName == "" {
		return nil, fmt.Errorf("versionId and name are required")
	}
	artifact := openapi.NewModelArtifactWithDefaults()
	artifact.SetName(artifactName)
	artifact.SetUri(uri)
	artifact.SetDescription(description)
	artifact.SetModelFormatName(modelFormatName)
	artifact.SetModelFormatVersion(modelFormatVersion)
	artifact.SetExternalID(externalId)

	a, resp, err := c.modelRegistryClient.ModelRegistryServiceAPI.CreateModelVersionArtifact(
		context.Background(), versionId,
	).
		Artifact(openapi.ModelArtifactAsArtifact(artifact)).Execute()
	if err != nil {
		if resp == nil {
			return nil, fmt.Errorf("error creating model version artifact: %w", err)
		} else {
			if resp.StatusCode != 201 {
				// TODO: Remove this workaround when model registry returns 404 when the model version is not found
				if isOpenAPIErrorOfKind(err, ErrVersionNotFound) {
					return nil, fmt.Errorf("%w. version id: %s", ErrVersionNotFound, versionId)
				}
				if isOpenAPIErrorOfKind(err, ErrAlreadyExists) {
					return nil, fmt.Errorf(
						"%w. version id: %s artifact name: %s", ErrArtifactExists, versionId, artifactName,
					)
				}
			}
			// This is a weird case where we got a response and an error that we're unable to handle.
			return nil, fmt.Errorf(
				"error while creating a model version artifact: server responded with %s %w", resp.Status, err,
			)
		}
	}
	if resp.StatusCode != 201 {
		return nil, fmt.Errorf("failed to create a model version artifact: %s", resp.Status)
	}
	return a, nil
}

// GetRegisteredModels returns a list of registered models
func (c *Client) GetRegisteredModels() ([]openapi.RegisteredModel, error) {
	models, resp, err := c.modelRegistryClient.ModelRegistryServiceAPI.GetRegisteredModels(
		context.Background(),
	).Execute()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to get registered models: %s", resp.Status)
	}
	return models.Items, nil
}

// ToMetadataValueMap converts a map of string to interface {} to a map of string to MetadataValue
func ToMetadataValueMap(
	metaData map[string]interface{},
) (map[string]openapi.MetadataValue, error) {
	props := make(map[string]openapi.MetadataValue)

	for k, v := range metaData {
		switch v := v.(type) {
		case string:
			mv := openapi.NewMetadataStringValueWithDefaults()
			mv.SetStringValue(v)
			props[k] = openapi.MetadataStringValueAsMetadataValue(
				mv,
			)
		case []interface{}:
			var ss []string
			for _, i := range v {
				if s, ok := i.(string); !ok {
					return nil, fmt.Errorf(
						"unsupported metadata value type for %s: %T. Only string and []string are supported", k, i,
					)
				} else {
					ss = append(ss, s)
				}
			}

			sv, err := encodeToBase64(StringList{Items: ss})
			if err != nil {
				return nil, fmt.Errorf("failed to encode metadata value for %s: %w", k, err)
			}
			mv := openapi.NewMetadataStructValueWithDefaults()
			mv.SetStructValue(sv)
			props[k] = openapi.MetadataStructValueAsMetadataValue(mv)
		default:
			continue
		}
	}
	return props, nil
}

func encodeToBase64(v interface{}) (string, error) {
	var buf bytes.Buffer
	encoder := base64.NewEncoder(base64.StdEncoding, &buf)
	err := json.NewEncoder(encoder).Encode(v)
	if err != nil {
		return "", err
	}
	encoder.Close()
	return buf.String(), nil
}
