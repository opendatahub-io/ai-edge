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
	"fmt"

	"github.com/spf13/cobra"

	"github.com/opendatahub-io/ai-edge/cli/pkg/commands/common"
	"github.com/opendatahub-io/ai-edge/cli/pkg/commands/flags"
)

var buildCmd = common.NewCmd(
	fmt.Sprintf(
		"build -%s model-id -%s version [-%s model-registry-url] [-%s namespace] [-%s kubeconfig]",
		flags.FlagModelID.Shorthand(),
		flags.FlagVersionName.Shorthand(),
		flags.FlagModelRegistryURL.Shorthand(),
		flags.FlagNamespace.Shorthand(),
		flags.FlagKubeconfig.Shorthand(),
	),
	"Build a synced edge model image",
	`Build a synced edge model image identified by the provided model image ID.

This command allows you to build an edge model image from the provided model image and model version using
the stored build parameters.
`,
	cobra.NoArgs,
	[]flags.Flag{
		flags.FlagModelID,
		flags.FlagVersionName,
		flags.FlagNamespace.SetParentFlag(),
		flags.FlagModelRegistryURL.SetParentFlag(),
		flags.FlagKubeconfig.SetParentFlag(),
	},
	common.SubCommandBuild,
	NewImagesModel,
)
