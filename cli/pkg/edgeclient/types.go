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

type modelImageStatus int

const (
	// ModelImageStatusUnknown - The status of the model image is unknown.
	ModelImageStatusUnknown modelImageStatus = iota
	// ModelImageStatusNeedsSync - The model image needs to be synced to the edge device.
	ModelImageStatusNeedsSync
	// ModelImageStatusSynced - The model image has been synced to the edge device.
	ModelImageStatusSynced
	// ModelImageStatusBuilding - The model image is being built.
	ModelImageStatusBuilding
	// ModelImageStatusLive - The model image is live on the container registry.
	ModelImageStatusLive
	// ModelImageStatusFailed - The model image build has failed.
	ModelImageStatusFailed
)

func (s modelImageStatus) String() string {
	return [...]string{"Unknown", "Needs Sync", "Synced", "Building", "Live", "Failed"}[s]
}

// Model - A registered model in the model registry.
type Model struct {
	ID          string
	Name        string
	Description string
}

// ModelImage - A model to be registered in the model registry and is suitable for deployment in edge environments.
type ModelImage struct {
	ID          string
	ModelID     string
	Name        string
	Description string
	Version     string
	URI         string
	NeedsSync   bool
	BuildParams map[string]interface{}
	Status      modelImageStatus
}
