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

package images

import (
	tea "github.com/charmbracelet/bubbletea"
	. "github.com/opendatahub-io/ai-edge/cli/pkg/commands/common"
	. "github.com/opendatahub-io/ai-edge/cli/pkg/commands/flags"
	"github.com/spf13/cobra"
)

var syncCmd = NewCmd(
	"sync <model-id> <model-version-name>",
	"Synchronize an edge model image information with the model registry",
	`Synchronize an edge model image information with the model registry.

This command allows you to synchronize an edge model image information with the model registry by ensuring that the
model image and the version artifact exist, are marked as edge compatible and have all the provided build parameters.

The command will return an error if the model image does not exist.

If the model version or the model version artifact do not exist, they will be created and if the model version is not
marked as edge compatible, it will be updated.

If the model version custom properties do not match the provided parameters, they will be updated.`,
	cobra.ExactArgs(2),
	[]Flag{
		FlagNamespace.SetInherited(), FlagModelRegistryUrl.SetInherited(), FlagKubeconfig.SetInherited(),
		FlagParams,
	},
	SubCommandSync,
	func(args []string, flags map[string]string, subCommand SubCommand) tea.Model {
		return NewImagesModel(
			args, flags, subCommand,
		)
	},
)

func init() {
	syncCmd.Flags().StringP("params", "p", "params.yaml", "Path to the build parameters file")
}
