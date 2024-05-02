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

package models

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/opendatahub-io/ai-edge/cli/pkg/commands/common"
	"github.com/opendatahub-io/ai-edge/cli/pkg/commands/flags"
)

var addCmd = common.NewCmd(
	fmt.Sprintf(
		"add -%s model-name -%s model-description [-%s version-name] [-%s model-registry-url] [-%s params-file]",
		flags.FlagModelName.Shorthand(),
		flags.FlagModelDescription.Shorthand(),
		flags.FlagVersionName.Shorthand(),
		flags.FlagModelRegistryURL.Shorthand(),
		flags.FlagParams.Shorthand(),
	),
	"Add model information to the model registry including the model name, model description, model version "+
		"and build parameters.",
	`Add model image information to the model registry including the model name, model description, model version and `+
		`build parameters.

If you don't provide a version name, the version name will be set to 'v1'.

Build parameters are provided via a YAML file with the following format:

params:
  - name: <parameter-name>
    value: <parameter-value>
  - name: <parameter-name>
    value: <parameter-value>
    ...
`,
	cobra.NoArgs,
	[]flags.Flag{
		flags.FlagModelName.SetRequired(),
		flags.FlagModelDescription.SetRequired(),
		flags.FlagVersionName,
		flags.FlagModelRegistryURL.SetParentFlag(),
		flags.FlagParams,
	},
	common.SubCommandAdd,
	NewTeaModel,
)
