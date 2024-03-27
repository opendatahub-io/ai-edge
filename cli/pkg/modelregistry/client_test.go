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
	"fmt"
	"net/http"
	"testing"

	"github.com/kubeflow/model-registry/pkg/openapi"
	"github.com/opendatahub-io/ai-edge/cli/pkg/httptest"
)

const (
	modelRegistryPath       = "/api/model_registry"
	getRegisteredModelsPath = "/v1alpha2/registered_models"
	success                 = "\u2713"
	failed                  = "\u2717"
)

func TestClient_GetRegisteredModels(t *testing.T) {

	table := []struct {
		name           string
		when           string
		responseBody   map[string][]map[string]interface{}
		expectedModels []openapi.RegisteredModel
	}{
		{
			name: "no models",
			when: "the model registry returns no models",
			responseBody: map[string][]map[string]interface{}{
				"items": {},
			},
			expectedModels: []openapi.RegisteredModel{},
		},
		{
			name: "one model",
			when: "the model registry returns one model",
			responseBody: map[string][]map[string]interface{}{
				"items": {
					convertModelToItem(newRegisteredModel("1", "model 1", "description 1")),
				},
			},
			expectedModels: []openapi.RegisteredModel{
				newRegisteredModel("1", "model 1", "description 1"),
			},
		},
		{
			name: "many models",
			when: "the model registry returns many models",
			responseBody: map[string][]map[string]interface{}{
				"items": {
					convertModelToItem(newRegisteredModel("1", "model 1", "description 1")),
					convertModelToItem(newRegisteredModel("2", "model 2", "description 2")),
				},
			},
			expectedModels: []openapi.RegisteredModel{
				newRegisteredModel("1", "model 1", "description 1"),
				newRegisteredModel("2", "model 2", "description 2"),
			},
		},
	}
	server := httptest.NewMockServer()
	server.Start()
	defer server.Close()
	t.Log("Given the need to test getting registered models from the model registry.")
	{
		for _, tt := range table {
			t.Run(
				tt.name, func(t *testing.T) {
					t.Logf("\t%s:\tWhen %s", t.Name(), tt.when)
					server.Reset()
					server.WithGet(
						fmt.Sprintf("%s%s", modelRegistryPath, getRegisteredModelsPath), httptest.Response{
							StatusCode:  http.StatusOK,
							ContentType: "application/json",
							Body:        tt.responseBody,
						},
					)
					client := NewClient(server.URL())
					models, err := client.GetRegisteredModels()
					if err != nil {
						t.Fatalf("\t%s:\t%s\tShould not receive an error. Got: %v", t.Name(), failed, err)
					}
					t.Logf("\t%s:\t%s\tShould not receive an error", t.Name(), success)

					if len(models) != len(tt.expectedModels) {
						t.Fatalf(
							"\t%s:\t%s\tShould receive %d models, got %d", t.Name(), failed, len(tt.expectedModels),
							len(models),
						)
						t.Logf("\t%s:\t%s\tShould receive %d models", t.Name(), success, len(tt.expectedModels))
					}

					for i, model := range models {
						compareRegisteredModels(t, model, tt.expectedModels[i])
					}
					t.Logf("\t%s:\t%s\tShould receive the expected models", t.Name(), success)
				},
			)
		}
	}
}

func compareRegisteredModels(t *testing.T, model openapi.RegisteredModel, expected openapi.RegisteredModel) {
	t.Helper()
	if model.GetId() != expected.GetId() {
		t.Fatalf("\t%s:\t%s\tShould receive the expected ID, got %s", t.Name(), failed, model.GetId())
	}
	if model.GetName() != expected.GetName() {
		t.Fatalf("\t%s:\t%s\tShould receive the expected Name, got %s", t.Name(), failed, model.GetName())
	}
	if model.GetDescription() != expected.GetDescription() {
		t.Fatalf("\t%s:\t%s\tShould receive the expected Description, got %s", t.Name(), failed, model.GetDescription())
	}
	// if model.GState != expected.State {
	// 	t.Fatalf("\t%s\tShould receive the expected State, got %s", failed, model.State)
	// }
	// t.Logf("\t%s\tShould receive the expected State", success)
	//
	// if model.CreateTimeSinceEpoch != expected.CreateTimeSinceEpoch {
	// 	t.Fatalf("\t%s\tShould receive the expected CreateTimeSinceEpoch, got %s", failed, model.CreateTimeSinceEpoch)
	// }
	// t.Logf("\t%s\tShould receive the expected CreateTimeSinceEpoch", success)
	//
	// if model.LastUpdateTimeSinceEpoch != expected.LastUpdateTimeSinceEpoch {
	// 	t.Fatalf(
	// 		"\t%s\tShould receive the expected LastUpdateTimeSinceEpoch, got %s", failed, model.LastUpdateTimeSinceEpoch,
	// 	)
	// }
	// t.Logf("\t%s\tShould receive the expected LastUpdateTimeSinceEpoch", success)
	//
	// if len(model.CustomProperties) != len(expected.CustomProperties) {
	// 	t.Fatalf("\t%s\tShould receive the expected CustomProperties, got %v", failed, model.CustomProperties)
	// }
	// t.Logf("\t%s\tShould receive the expected CustomProperties", success)
}

func convertModelToItem(model openapi.RegisteredModel) map[string]interface{} {
	return map[string]interface{}{
		"id":                       model.GetId(),
		"name":                     model.GetName(),
		"description":              model.GetDescription(),
		"customProperties":         model.GetCustomProperties(),
		"state":                    model.GetState(),
		"externalID":               model.GetExternalID(),
		"createTimeSinceEpoch":     model.GetCreateTimeSinceEpoch(),
		"lastUpdateTimeSinceEpoch": model.GetLastUpdateTimeSinceEpoch(),
	}
}

func newRegisteredModel(id, name, description string) openapi.RegisteredModel {
	model := openapi.NewRegisteredModel()
	model.SetId(id)
	model.SetName(name)
	model.SetDescription(description)
	return *model
}

func TestClient(t *testing.T) {
	modelRegistryURL := "http://localhost:8080"
	client := NewClient(modelRegistryURL)
	// assert.Equal(t, modelRegistryURL, client.modelRegistryURL)
	// assert.Assert(t, client.modelRegistryClient != nil)

	// _, err := client.GetModelVersionArtifacts("123")
	// v, err := client.CreateRegisteredModel("tensorflow-housing-test", "tensorflow-housing-test - ", nil)
	s := openapi.NewMetadataStringValueWithDefaults()
	s.SetStringValue("disabled")

	v, err := client.UpdateModelVersion(
		"6", map[string]openapi.MetadataValue{
			"edge":   openapi.MetadataStringValueAsMetadataValue(s),
			"deploy": openapi.MetadataStringValueAsMetadataValue(s),
		},
	)
	// v, err := client.FindModelVersionArtifact("4", "tensorflow-housing-6")
	// v, err := client.CreateModelArtifact(
	// 	"6", "tensorflow-housing-test", "tensorflow-housing-test - ", "onnx", "1", "aws-connection-mybucket",
	// )
	fmt.Printf("v %v, err %v\n", v.GetCustomProperties(), err)
	// assert.NilError(t, err)
	// assert.Assert(t, errors.Is(err, ErrVersionNotFound))
	// assert.Equal(t, "123", vs)
}
