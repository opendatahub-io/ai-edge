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

package edgeclient

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/opendatahub-io/ai-edge/cli/pkg/modelregistry"
)

// Client is a client representing the edge environment
//
// This client can be used to create and manage models and model container images suitable for deployment in edge environments.
type Client struct {
	modelRegistryClient *modelregistry.Client
}

// NewClient creates a new Client to interact with the edge environment. It requires the URL of the model registry service.
//
// This client can be used to create and manage models and model container images suitable for deployment in edge environments.
func NewClient(modelRegistryURL string) *Client {
	return &Client{
		modelRegistryClient: modelregistry.NewClient(modelRegistryURL),
	}
}

// GetModels returns a list of models in the model registry.
func (c *Client) GetModels() ([]Model, error) {
	models, err := c.modelRegistryClient.GetRegisteredModels()
	if err != nil {
		return nil, fmt.Errorf("failed to get models: %w", err)
	}
	ms := make([]Model, len(models))
	for i, m := range models {
		ms[i] = Model{
			Id:          m.GetId(),
			Name:        m.GetName(),
			Description: m.GetDescription(),
		}
	}
	return ms, nil
}

// AddNewModelWithImage adds a model to the model registry along with the model version and all build parameters required to
// create a model container image using a tekton pipeline run.
func (c *Client) AddNewModelWithImage(
	modelName string,
	modelDescription string,
	modelVersion string,
	uri string,
	parameters map[string]interface{},
) (*ModelImage, error) {

	if modelName == "" || modelDescription == "" || modelVersion == "" {
		return nil, fmt.Errorf("model name, description, version, and URI are required")
	}

	modelFormatName := "ContainerImage"
	// This will be used to flag the model as edge compatible (i.e. has the required metadata to be built by the edge pipeline)
	parameters["edgeCompatible"] = "true"
	externalId := getImageId(modelName, modelVersion, modelName)
	md, err := modelregistry.ToMetadataValueMap(parameters)
	if err != nil {
		return nil, fmt.Errorf("failed to add model image: %w", err)
	}
	m, v, a, err := c.modelRegistryClient.AutoRegisterModelVersionArtifact(
		modelName, modelDescription, modelVersion, modelName, externalId, uri, modelFormatName, "", md,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to add model image: %w", err)
	}

	return &ModelImage{
		Id:          a.ModelArtifact.GetExternalID(),
		ModelId:     m.GetId(),
		Name:        m.GetName(),
		Description: m.GetDescription(),
		Version:     v.GetName(),
		URI:         a.ModelArtifact.GetUri(),
	}, nil
}

func getImageId(registeredModelName, modelVersionName, artifactName string) string {
	return shortHash(fmt.Sprintf("%s:%s:%s", registeredModelName, modelVersionName, artifactName))
}

func shortHash(s string) string {
	hasher := sha256.New()
	hasher.Write([]byte(s))
	hash := hex.EncodeToString(hasher.Sum(nil))
	return hash[:4]
}
