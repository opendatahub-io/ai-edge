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

package common

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/opendatahub-io/ai-edge/cli/pkg/commands/flags"
	"github.com/spf13/cobra"
)

type SubCommand int

const (
	SubCommandList SubCommand = iota
	SubCommandSync
	SubCommandBuild
	SubCommandAdd
)

func NewCmd(
	use, short, long string,
	args cobra.PositionalArgs,
	flags []flags.Flag,
	command SubCommand,
	modelFactory func(args []string, flags map[string]string, subCommand SubCommand) tea.Model,
) *cobra.Command {

	cmd := cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		Args:  args,
		Run: func(cmd *cobra.Command, args []string) {
			ff := make(map[string]string)
			for _, f := range flags {
				v := ""
				err := error(nil)
				if f.IsInherited() {
					v, err = cmd.InheritedFlags().GetString(f.String())
					if err != nil {
						cmd.PrintErrf("Error reading inherited flag %s: %v\n", f, err)
						os.Exit(1)
					}
				} else {
					v, err = cmd.Flags().GetString(f.String())
					if err != nil {
						cmd.PrintErrf("Error reading flag %s: %v\n", f, err)
						os.Exit(1)
					}
				}
				ff[f.String()] = v
			}
			_, err := tea.NewProgram(modelFactory(args, ff, command)).Run()
			if err != nil {
				cmd.PrintErrf("Error: %v\n", err)
				os.Exit(1)
			}
		},
	}

	return &cmd
}
