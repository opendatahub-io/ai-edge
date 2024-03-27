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

type ModelImageStatus int

const (
	ModelImageStatusUnknown ModelImageStatus = iota
	ModelImageNeedsSync
	ModelImageStatusSynced
	ModelImageStatusBuilding
	ModelImageStatusLive
	ModelImageStatusFailed
)

func (s ModelImageStatus) String() string {
	return [...]string{"Unknown", "Needs Sync", "Synced", "Building", "Live", "Failed"}[s]
}

type Model struct {
	Id          string
	Name        string
	Description string
}

// ModelImage - A model to be registered in the model registry and is suitable for deployment in edge environments.
type ModelImage struct {
	Id          string
	ModelId     string
	Name        string
	Description string
	Version     string
	URI         string
	NeedsSync   bool
	BuildParams map[string]interface{}
	Status      ModelImageStatus
}

type PipelineRun struct {
	Name      string
	Namespace string
}
