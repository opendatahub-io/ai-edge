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
	"strings"

	"github.com/opendatahub-io/ai-edge/cli/pkg/modelregistry"
)

// Client is a client representing the edge environment
//
// This client can be used to create and manage models and model container images suitable for deployment in edge environments.
type Client struct {
	modelRegistryClient *modelregistry.Client
	kubeconfig          string
}

// NewClient creates a new Client to interact with the edge environment. It requires the URL of the model registry service.
//
// This client can be used to create and manage models and model container images suitable for deployment in edge environments.
func NewClient(modelRegistryURL, kubeconfig string) *Client {
	if !strings.Contains(modelRegistryURL, "://") {
		modelRegistryURL = "http://" + modelRegistryURL
	}
	return &Client{
		modelRegistryClient: modelregistry.NewClient(modelRegistryURL),
		kubeconfig:          kubeconfig,
	}
}
