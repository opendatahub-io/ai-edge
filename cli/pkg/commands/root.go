/*
Copyright © 2024 Open Data Hub Authors

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

package commands

import (
	"os"

	"github.com/opendatahub-io/ai-edge/cli/pkg/commands/flags"
	"github.com/opendatahub-io/ai-edge/cli/pkg/commands/images"
	"github.com/opendatahub-io/ai-edge/cli/pkg/commands/models"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "odh-cli",
	Short: "Manage Open Data Hub resources from the command line.",
	Long: `Manage Open Data Hub resources from the command line.

This application is a tool to perform various operations on Open Data Hub.`,

}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	for _, f := range flags.Flags {
		if f.IsRootFlag() {
			rootCmd.PersistentFlags().StringP(f.String(), f.Shorthand(), f.Value(), f.Usage())
		}
	}
	rootCmd.AddCommand(images.Cmd)
	rootCmd.AddCommand(models.Cmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
}
