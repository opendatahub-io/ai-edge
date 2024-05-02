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

	"github.com/opendatahub-io/ai-edge/cli/pkg/commands/common"
	. "github.com/opendatahub-io/ai-edge/cli/pkg/commands/common"
	"github.com/opendatahub-io/ai-edge/cli/pkg/commands/flags"
	. "github.com/opendatahub-io/ai-edge/cli/pkg/commands/flags"
	"github.com/spf13/cobra"
)

var describeCmd = common.NewCmd(
	fmt.Sprintf(
		"describe -%s model-id -%s version [-%s model-registry-url]",
		flags.FlagModelID.Shorthand(),
		flags.FlagVersionName.Shorthand(),
		flags.FlagModelRegistryURL.Shorthand(),
	),
	"View details of an edge model image.",
	`View details of an edge model image.

This command allows you to view details of a specific edge model image along with its parameters.
`,
	cobra.NoArgs,
	[]Flag{
		flags.FlagModelID.SetRequired(),
		flags.FlagVersionName.SetRequired(),
		flags.FlagModelRegistryURL.SetParentFlag(),
	},
	SubCommandDescribe,
	NewImagesModel,
)
