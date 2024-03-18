/*
Copyright 2023 KStreamer Authors

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
	"context"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	openapiclient "github.com/kubeflow/model-registry/pkg/openapi"
	"github.com/spf13/cobra"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("#04B575"))

var errorStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FF0000")).
	Bold(true).
	Height(4).
	Width(120)

type registeredModelsMsg *openapiclient.RegisteredModelList

type errMsg struct{ err error }

func (e errMsg) Error() string { return e.err.Error() }

func listRegisteredModels() tea.Msg {
	configuration := openapiclient.NewConfiguration()
	configuration.Servers = openapiclient.ServerConfigurations{
		{
			URL: modelRegistryURL,
		},
	}

	apiClient := openapiclient.NewAPIClient(configuration)
	models, httpRes, err := apiClient.ModelRegistryServiceAPI.GetRegisteredModels(
		context.Background(),
	).Execute()
	if err != nil {
		return errMsg{err}
	}
	if httpRes.StatusCode != 200 {
		fmt.Printf("not 200: %d\n", httpRes.StatusCode)
		return errMsg{
			fmt.Errorf(
				"Failed to get models, calling the model registry API returned status code %d\n", httpRes.StatusCode,
			),
		}
	}
	return registeredModelsMsg(models)
}

type model struct {
	registeredModelsList *openapiclient.RegisteredModelList
	cursor               int
	choices              int
	err                  error
}

func (m model) Init() tea.Cmd {
	return listRegisteredModels
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case errMsg:
		m.err = msg
		return m, tea.Quit

	case registeredModelsMsg:
		m.registeredModelsList = msg
		m.choices = len(m.registeredModelsList.Items)
		return m, tea.Quit
	}
	return m, nil
}

func (m model) View() string {
	if m.err != nil {
		return errorStyle.Render(fmt.Sprintf("Error: %s", m.err))
		// return ""
	}
	columns := []table.Column{
		{Title: "Id", Width: 4},
		{Title: "Name", Width: 20},
		{Title: "Description", Width: 60},
	}

	rows := make([]table.Row, 0)

	if m.registeredModelsList != nil {
		for _, model := range m.registeredModelsList.Items {
			rows = append(
				rows, table.Row{
					model.GetId(),
					model.GetName(),
					model.GetDescription(),
				},
			)
		}
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithHeight(len(rows)+1),
	)

	s := table.DefaultStyles()
	s.Cell.Foreground(lipgloss.Color("#FFF"))
	s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#04B575")).
		BorderBottom(true).
		Bold(true)
	t.SetStyles(s)
	t.SetCursor(m.cursor)
	return baseStyle.Render(t.View()) + "\n"
}

// modelsCmd represents the models command
var modelsCmd = &cobra.Command{
	Use:   "models",
	Short: "Manage models",
	Long: `Manage Open Data Hub models from the command line.

This command will list all the registered models available in the Open Data Hub model registry.`,
	Run: func(cmd *cobra.Command, args []string) {
		m, err := tea.NewProgram(model{}).Run()
		if err != nil {
			cmd.PrintErrf("Error: %v\n", err)
			os.Exit(1)
		}
		if m.(model).err != nil {
			// cmd.PrintErrf("Error: %v\n", m.(model).err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(modelsCmd)
}
