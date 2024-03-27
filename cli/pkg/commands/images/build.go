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

var buildCmd = NewCmd(
	"build <id>",
	"Build a synced edge model image",
	`Build a synced edge model image identified by the provided model image ID.

This command allows you to build an edge model image from the provided model image which points to a model version using
the stored build parameters.

This command will fail if the model image needs sync (using the sync command).`,
	cobra.ExactArgs(1),
	[]Flag{
		FlagNamespace.SetInherited(), FlagModelRegistryUrl.SetInherited(), FlagKubeconfig.SetInherited(),
		FlagParams,
	},
	SubCommandBuild,
	func(args []string, flags map[string]string, subCommand SubCommand) tea.Model {
		return NewImagesModel(
			args, flags, subCommand,
		)
	},
)

func init() {
	buildCmd.Flags().StringP("params", "p", "params.yaml", "Path to the build parameters file")
}
