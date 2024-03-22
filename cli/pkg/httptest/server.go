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

package httptest

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
)

type Response struct {
	StatusCode  int
	ContentType string
	Body        any
}

type MockServer struct {
	server *httptest.Server
	routes map[string]map[string]Response
}

func NewMockServer() *MockServer {

	return &MockServer{
		routes: map[string]map[string]Response{
			"GET":    {},
			"POST":   {},
			"PUT":    {},
			"DELETE": {},
			"PATCH":  {},
		},
	}
}

func (m *MockServer) WithGet(path string, response Response) {
	m.routes["GET"][path] = response
}

func (m *MockServer) WithPost(path string, response Response) {
	m.routes["POST"][path] = response
}

func (m *MockServer) Reset() {
	m.routes = map[string]map[string]Response{
		"GET":    {},
		"POST":   {},
		"PUT":    {},
		"DELETE": {},
		"PATCH":  {},
	}
}

func (m *MockServer) Start() {
	// Create a new httptest server using the handler
	m.server = httptest.NewServer(m.handler())
}

func (m *MockServer) handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the response for the request
		response, ok := m.routes[r.Method][r.URL.Path]
		if !ok {
			http.NotFound(w, r)
			return
		}

		// Set content type
		w.Header().Set("Content-Type", response.ContentType)

		// Set the status code
		w.WriteHeader(response.StatusCode)

		// Write the response
		if err := json.NewEncoder(w).Encode(response.Body); err != nil {
			log.Fatalf("Error encoding response: %v", err)
		}
	}
}

func (m *MockServer) Close() {
	m.server.Close()
}

func (m *MockServer) URL() string {
	return m.server.URL
}

// func (m *MockServer)

// 	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// Check for the correct path
// 		if r.URL.Path != "/api/model_registry/v1alpha2/registered_models" {
// 			http.NotFound(w, r)
// 			return
// 		}
//
// 		// Define the response
// 		response := map[string][]map[string]interface{}{
// 			"items": {
// 				{
// 					"customProperties": map[string]map[string]string{
// 						"additionalProp1": {"int_value": "string"},
// 						"additionalProp2": {"int_value": "string"},
// 						"additionalProp3": {"int_value": "string"},
// 					},
// 					"description":               "string",
// 					"externalID":                "string",
// 					"name":                      "string",
// 					"id":                        "string",
// 					"createTimeSinceEpoch":      "string",
// 					"lastUpdateTimeSinceEpoch": "string",
// 					"state":                     "LIVE",
// 				},
// 			},
// 		}
//
// 		// Set content type
// 		w.Header().Set("Content-Type", "application/json")
//
// 		// Write the response
// 		if err := json.NewEncoder(w).Encode(response); err != nil {
// 			log.Fatalf("Error encoding response: %v", err)
// 		}
// 	})
//
// 	// Create a new httptest server using the handler
// 	server := httptest.NewServer(handler)
// 	defer server.Close()
//
// 	// Use server.URL to access the server in your tests
// 	log.Println("Mock server running at:", server.URL)
//
// 	// In a test, you can make HTTP requests to server.URL
// 	// For example: http.Get(server.URL + "/api/model_registry/v1alpha2/registered_models")
// }
//
