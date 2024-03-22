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

// SubCommand is a type to represent the subcommand
type SubCommand int

const (
	// SubCommandList is a subcommand to list items
	SubCommandList SubCommand = iota
	// SubCommandAdd is a subcommand to add items
	SubCommandAdd
	// SubCommandUpdate is a subcommand to sync items
	SubCommandUpdate
	// SubCommandBuild is a subcommand to build items
	SubCommandBuild
	// SubCommandDescribe is a subcommand to view details
	SubCommandDescribe
)

// NewCmd creates a new cobra command.
//
// The command will create a new tea program, passing the model created by the modelFactory, and run it.
// The modelFactory will be called with the args, flags and subCommand.
//
// Example:
//
//	cmd := NewCmd(
//		"images",
//		"List images",
//		`List images`,
//		cobra.ExactArgs(3),
//		[]flags.Flag{flags.FlagModelRegistryUrl},
//		SubCommandList,
//		func(args []string, flags map[string]string, subCommand SubCommand) tea.Model {
//			return NewImagesModel(args, flags, subCommand)
//		},
//	)
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
				if f.IsParentFlag() {
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

	// Disable the addition of [flags] to the usage line of a command when printing help or generating docs
	cmd.DisableFlagsInUseLine = true

	cmd.Flags().SortFlags = false

	for _, f := range flags {
		if !f.IsParentFlag() {
			if f.IsInherited() {
				cmd.PersistentFlags().StringP(f.String(), f.Shorthand(), f.Value(), f.Usage())
			} else {
				cmd.Flags().StringP(f.String(), f.Shorthand(), f.Value(), f.Usage())
			}
			if f.IsRequired() {
				err := cmd.MarkFlagRequired(f.String())
				if err != nil {
					cmd.PrintErrf("Error marking flag %s as required: %v\n", f, err)
					os.Exit(1)
				}
			}
		}
	}

	return &cmd
}
