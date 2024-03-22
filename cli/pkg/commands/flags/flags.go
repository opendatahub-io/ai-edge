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

package flags

import (
	"fmt"
	"os"
)

// Flag represents a command line flag.
//
// Flags can be inherited by subcommands, in which case they will be passed to the subcommand.
type Flag struct {
	name       string
	inherited  bool // Flag is inherited by subcommands
	parentFlag bool // Flag is defined in the parent command
	shorthand  string
	required   bool
	value      string // Default value, only if the flag is required
	usage      string
}

var (
	// FlagModelRegistryURL is the URL of the model registry
	FlagModelRegistryURL = Flag{
		name: "model-registry-url", inherited: true, shorthand: "r", value: "http://localhost:8080",
		usage: "URL of the model registry",
	}

	// FlagKubeconfig is the path to the kubeconfig file
	FlagKubeconfig = Flag{
		name: "kubeconfig", inherited: true, shorthand: "k",
		value: fmt.Sprintf("%s/.kube/config", os.Getenv("HOME")),
		usage: "path to the kubeconfig file",
	}

	// FlagNamespace is the namespace to use
	FlagNamespace = Flag{
		name:      "namespace",
		shorthand: "n",
		value:     "default",
		usage:     "namespace to use",
	}

	// FlagParams is the path to the build parameters file
	FlagParams = Flag{
		name:      "params",
		shorthand: "p",
		value:     "params.yaml",
		usage:     "path to the build parameters file",
	}

	// FlagModelID is the ID of the model
	FlagModelID = Flag{
		name:      "model-id",
		shorthand: "i",
		usage:     "ID of the model",
	}

	// FlagModelName is the name of the model
	FlagModelName = Flag{
		name:      "model-name",
		shorthand: "m",
		usage:     "model name",
	}

	// FlagModelDescription is the description of the model
	FlagModelDescription = Flag{
		name:      "model-description",
		shorthand: "d",
		usage:     "model description",
	}

	// FlagVersionName is the name of the model version
	FlagVersionName = Flag{
		name:      "version-name",
		shorthand: "v",
		value:     "v1",
		usage:     "model version name. A model can have multiple versions.",
	}

	// FlagModelImageID is the ID of the model image
	FlagModelImageID = Flag{
		name:      "image-id",
		shorthand: "g",
		usage:     "model image ID",
	}
)

// String returns the name of the flag.
func (f Flag) String() string {
	return f.name
}

// SetInherited sets the flag to be inherited by subcommands.
func (f Flag) SetInherited() Flag {
	f.inherited = true
	return f
}

// IsInherited returns true if the flag is inherited by subcommands.
func (f Flag) IsInherited() bool {
	return f.inherited
}

// SetParentFlag sets the flag as one that's defined in the parent command.
func (f Flag) SetParentFlag() Flag {
	f.parentFlag = true
	return f
}

// IsParentFlag returns true if the flag is defined in the parent command.
func (f Flag) IsParentFlag() bool {
	return f.parentFlag
}

// SetRequired sets the flag as required.
func (f Flag) SetRequired() Flag {
	f.required = true
	f.usage += " (required)"
	return f
}

// IsRequired returns true if the flag is required.
func (f Flag) IsRequired() bool {
	return f.required
}

// Shorthand returns the shorthand of the flag.
func (f Flag) Shorthand() string {
	return f.shorthand
}

// Value returns the value of the flag.
func (f Flag) Value() string {
	return f.value
}

// Usage returns the usage of the flag.
func (f Flag) Usage() string {
	return f.usage
}
