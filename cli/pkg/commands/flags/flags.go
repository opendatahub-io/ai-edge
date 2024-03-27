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

type Flag struct {
	name string
	inherited bool
	rootFlag bool
	shorthand string
	value string
	usage string
}

var (
	FlagModelRegistryUrl = Flag{
		name: "model-registry-url", rootFlag: true, shorthand: "m", value: "http://localhost:8080",
		usage: "URL of the model registry",
	}
	FlagKubeconfig       = Flag{
		name: "kubeconfig", rootFlag: true, shorthand: "k", value: fmt.Sprintf("%s/.kube/config", os.Getenv("HOME")),
		usage: "Path to the kubeconfig file",
	}
	FlagNamespace        = Flag{name: "namespace"}
	FlagParams           = Flag{name: "params"}

	Flags = []Flag{FlagKubeconfig, FlagModelRegistryUrl, FlagNamespace, FlagParams}
)

func (f Flag) String() string {
	return f.name
}

func (f Flag) SetInherited() Flag {
	f.inherited = true
	return f
}

func (f Flag) IsInherited() bool {
	return f.inherited
}

func (f Flag) IsRootFlag() bool {
	return f.rootFlag
}

func (f Flag) Shorthand() string {
	return f.shorthand
}

func (f Flag) Value() string {
	return f.value
}

func (f Flag) Usage() string {
	return f.usage
}
