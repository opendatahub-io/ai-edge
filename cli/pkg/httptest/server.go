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

// Response represents a response from the mock server that can be set for a given path and method
type Response struct {
	StatusCode  int
	ContentType string
	Body        any
}

// MockServer is a simple mock server that can be used to mock HTTP responses
type MockServer struct {
	server *httptest.Server
	routes map[string]map[string]Response
}

// NewMockServer creates a new mock server
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

// WithGet sets a response for a GET request to the given path
func (m *MockServer) WithGet(path string, response Response) {
	m.routes["GET"][path] = response
}

// WithPost sets a response for a POST request to the given path
func (m *MockServer) WithPost(path string, response Response) {
	m.routes["POST"][path] = response
}

// Reset clears all the set responses
func (m *MockServer) Reset() {
	m.routes = map[string]map[string]Response{
		"GET":    {},
		"POST":   {},
		"PUT":    {},
		"DELETE": {},
		"PATCH":  {},
	}
}

// Start starts the mock server
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

// Close closes the mock server
func (m *MockServer) Close() {
	m.server.Close()
}

// GetURL returns the URL of the mock server
func (m *MockServer) GetURL() string {
	return m.server.URL
}
