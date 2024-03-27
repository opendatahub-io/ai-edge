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
	"errors"
	"fmt"
	"strings"

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

func (c *Client) GetRegisteredModelById(id string) (*openapi.RegisteredModel, error) {

	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	model, resp, err := c.modelRegistryClient.ModelRegistryServiceAPI.GetRegisteredModel(
		context.Background(), id,
	).Execute()

	if err != nil {
		if resp == nil {
			return nil, fmt.Errorf("error getting registered model: %w", err)
		} else {
			// Currently model registry returns 500 when the model is not found. This is a workaround to handle the
			// error until the model registry is fixed. The workaround is to check the error message and return
			// ErrModelNotFound if the error message contains the expected error message.
			// TODO: Remove this workaround when model registry returns 404 when the model is not found
			if resp.StatusCode != 200 && isOpenAPIErrorOfKind(err, ErrModelNotFound) {
				return nil, fmt.Errorf("%w1d: %s", ErrModelNotFound, id)
			}
			// This is a weird case where we got a response and an error that we're unable to handle.
			return nil, fmt.Errorf(
				"error while getting registered model: server responded with %s %w", resp.Status, err,
			)
		}
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to get registered model: %s", resp.Status)
	}
	return model, nil
}

// GetModelVersions returns a list of model versions for a given model identified by modelId
func (c *Client) GetModelVersions(modelId string) ([]openapi.ModelVersion, error) {

	if modelId == "" {
		return nil, fmt.Errorf("registeredModelId is required")
	}

	versions, resp, err := c.modelRegistryClient.ModelRegistryServiceAPI.GetRegisteredModelVersions(
		context.Background(),
		modelId,
	).Execute()
	if err != nil {
		if resp == nil {
			return nil, fmt.Errorf("error getting model versions: %w", err)
		} else {
			// Currently model registry returns 500 when the model is not found. This is a workaround to handle the
			// error until the model registry is fixed. The workaround is to check the error message and return
			// ErrModelNotFound if the error message contains the expected error message.
			// TODO: Remove this workaround when model registry returns 404 when the model is not found
			if resp.StatusCode != 200 && isOpenAPIErrorOfKind(err, ErrModelNotFound) {
				return nil, fmt.Errorf("%w. model id: %s", ErrModelNotFound, modelId)
			}
			// This is a weird case where we got a response and an error that we're unable to handle.
			return nil, fmt.Errorf("error while getting model versions: server responded with %s %w", resp.Status, err)
		}
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to get model versions: %s", resp.Status)
	}
	return versions.Items, nil
}

// GetModelVersionById returns a model version by its ID
func (c *Client) GetModelVersionById(versionId string) (*openapi.ModelVersion, error) {

	if versionId == "" {
		return nil, fmt.Errorf("id is required")
	}

	version, resp, err := c.modelRegistryClient.ModelRegistryServiceAPI.GetModelVersion(
		context.Background(), versionId,
	).Execute()
	if err != nil {
		if resp == nil {
			return nil, fmt.Errorf("error getting model version: %w", err)
		} else {
			// Currently model registry returns 500 when the model version is not found. This is a workaround to handle
			// the error until the model registry is fixed. The workaround is to check the error message and return
			// ErrVersionNotFound if the error message contains the expected error message.
			// TODO: Remove this workaround when model registry returns 404 when the model version is not found
			if resp.StatusCode != 200 && isOpenAPIErrorOfKind(err, ErrVersionNotFound) {
				return nil, fmt.Errorf("%w. version id: %s", ErrVersionNotFound, versionId)
			}
			// This is a weird case where we got a response and an error that we're unable to handle.
			return nil, fmt.Errorf("error while getting model version: server responded with %s %w", resp.Status, err)
		}
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to get model version: %s", resp.Status)
	}
	return version, nil
}

// GetModelVersionArtifacts returns a list of model version artifacts for a given model version identified by versionId
func (c *Client) GetModelVersionArtifacts(versionId string) ([]openapi.Artifact, error) {

	if versionId == "" {
		return nil, fmt.Errorf("modelVersionId is required")

	}
	artifacts, resp, err := c.modelRegistryClient.ModelRegistryServiceAPI.GetModelVersionArtifacts(
		context.Background(),
		versionId,
	).Execute()
	if err != nil {
		if resp == nil {
			return nil, fmt.Errorf("error getting model version artifacts: %w", err)
		} else {
			// Currently model registry returns 500 when the model version is not found. This is a workaround to handle
			// the error until the model registry is fixed. The workaround is to check the error message and return
			// ErrVersionNotFound if the error message contains the expected error message.
			// TODO: Remove this workaround when model registry returns 404 when the model version is not found
			if resp.StatusCode != 200 && isOpenAPIErrorOfKind(err, ErrVersionNotFound) {
				return nil, fmt.Errorf("%w. version id: %s", ErrVersionNotFound, versionId)
			}
			// This is a weird case where we got a response and an error that we're unable to handle.
			return nil, fmt.Errorf(
				"error while getting model version artifacts: server responded with %s %w", resp.Status, err,
			)
		}
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to get model version artifacts: %s", resp.Status)
	}
	return artifacts.Items, nil
}

// FindModelVersion returns a model version by its name for a given model identified by modelId
//
// Errors:
// - ErrFindModelVersion is returned when no model version is found
func (c *Client) FindModelVersion(modelId, versionName string) (*openapi.ModelVersion, error) {
	if versionName == "" {
		return nil, fmt.Errorf("versionName is required")
	}

	v, r, err := c.modelRegistryClient.ModelRegistryServiceAPI.FindModelVersion(context.Background()).
		Name(versionName).ParentResourceID(modelId).Execute()
	if err != nil {
		if r == nil {
			return nil, fmt.Errorf("error finding model version by name: %w", err)
		} else {
			if r.StatusCode != 200 && isOpenAPIErrorOfKind(err, ErrFindModelVersion) {
				return nil, fmt.Errorf("%w. version name: %s", ErrFindModelVersion, versionName)
			}
			return nil, fmt.Errorf("error finding model version by name: server responded with %v %w", r.Status, err)
		}
	}
	if r.StatusCode != 200 {
		return nil, fmt.Errorf("failed to find model version by name: %s", r.Status)
	}
	return v, nil
}

// FindModelVersionArtifact returns a model version artifact by its name for a given model version identified by versionId
//
// Errors:
// - ErrFindArtifact is returned when no model version artifact is found
func (c *Client) FindModelVersionArtifact(versionId, artifactName string) (*openapi.ModelArtifact, error) {
	if artifactName == "" {
		return nil, fmt.Errorf("artifactName is required")
	}

	aa, err := c.GetModelVersionArtifacts(versionId)
	if err != nil {
		return nil, err
	}
	for _, a := range aa {
		if a.ModelArtifact != nil && a.ModelArtifact.GetName() == artifactName {
			return a.ModelArtifact, nil
		}
	}
	return nil, fmt.Errorf(
		"%w. Failed to find model version artifact by name: %s version id: %s", ErrFindArtifact, artifactName,
		versionId,
	)
}

// UpdateModelVersion updates a model version identified by versionId
//
// Errors:
// - ErrVersionNotFound is returned when the model version is not found
func (c *Client) UpdateModelVersion(
	versionId string,
	newMetaData map[string]openapi.MetadataValue,
) (*openapi.ModelVersion, error) {
	if versionId == "" {
		return nil, fmt.Errorf("versionId is required")
	}
	v, err := c.GetModelVersionById(versionId)
	if err != nil {
		if errors.Is(err, ErrVersionNotFound) {
			return nil, err
		}
		return nil, fmt.Errorf("error getting model version: %w", err)
	}
	if newMetaData == nil {
		return v, nil
	}
	v.SetCustomProperties(newMetaData)
	v, resp, err := c.modelRegistryClient.ModelRegistryServiceAPI.UpdateModelVersion(context.Background(), versionId).
		ModelVersion(*v).Execute()
	if err != nil {
		return nil, fmt.Errorf("error updating model version: %w", err)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to update model version: %s", resp.Status)
	}
	return v, nil
}

// UpdateModelArtifact updates a model version artifact identified by artifactId
//
// Errors:
// - ErrArtifactNotFound is returned when the model version artifact is not found
func (c *Client) UpdateModelArtifact(
	artifactId string,
	externalId string,
) (*openapi.ModelArtifact, error) {
	if artifactId == "" {
		return nil, fmt.Errorf("artifactId is required")

	}
	artifactUpdate := openapi.NewModelArtifactUpdate()
	artifactUpdate.SetExternalID(externalId)
	a, resp, err := c.modelRegistryClient.ModelRegistryServiceAPI.UpdateModelArtifact(
		context.Background(), artifactId,
	).ModelArtifactUpdate(
		*artifactUpdate,
	).Execute()
	if err != nil {
		if resp == nil {
			return nil, fmt.Errorf("error updating model version artifact: %w", err)
		} else {
			if resp.StatusCode != 200 && isOpenAPIErrorOfKind(err, ErrArtifactNotFound) {
				return nil, fmt.Errorf("%w. artifact id: %s", ErrArtifactNotFound, artifactId)
			}
			return nil, fmt.Errorf(
				"error updating model version artifact: server responded with %s %w", resp.Status, err,
			)
		}
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to update model version artifact: %s", resp.Status)
	}
	return a, nil
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

// FromMetadataValueMap converts a map of MetadataValue to a map of string to interface {}
func FromMetadataValueMap(
	metaData map[string]openapi.MetadataValue,
) (map[string]interface{}, error) {
	props := make(map[string]interface{})
	for k, v := range metaData {
		switch {
		case v.MetadataStringValue != nil:
			props[k] = v.MetadataStringValue.GetStringValue()
		case v.MetadataStructValue != nil:
			var sl StringList
			if err := decodeFromBase64(&sl, v.MetadataStructValue.GetStructValue()); err != nil {
				return nil, fmt.Errorf("failed to decode metadata value for %s: %w", k, err)
			}
			props[k] = sl.Items
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

func decodeFromBase64(v interface{}, enc string) error {
	decoder := base64.NewDecoder(base64.StdEncoding, strings.NewReader(enc))
	return json.NewDecoder(decoder).Decode(v)
	// return json.NewDecoder(base64.NewDecoder(base64.StdEncoding, strings.NewReader(enc))).Decode(v)
}
