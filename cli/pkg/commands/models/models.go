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

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"

	"github.com/opendatahub-io/ai-edge/cli/pkg/commands/common"
	"github.com/opendatahub-io/ai-edge/cli/pkg/commands/flags"
	"github.com/opendatahub-io/ai-edge/cli/pkg/edgeclient"
	"github.com/opendatahub-io/ai-edge/cli/pkg/pipelines"
)

type registeredModelsMsg []edgeclient.Model
type newModelAddedMsg struct{}

type teaModel struct {
	args                 []string
	flags                map[string]string
	edgeClient           *edgeclient.Client
	registeredModelsList []edgeclient.Model
	err                  error
	subCommand           common.SubCommand
}

// NewTeaModel creates a new bubbletea model for the models command
func NewTeaModel(args []string, flgs map[string]string, subCommand common.SubCommand) tea.Model {
	return &teaModel{
		args:       args,
		flags:      flgs,
		edgeClient: edgeclient.NewClient(flgs[flags.FlagModelRegistryURL.String()]),
		subCommand: subCommand,
	}
}

// Init initializes the model according to the subcommand
func (m teaModel) Init() tea.Cmd {
	switch m.subCommand {
	case common.SubCommandList:
		return m.listRegisteredModels()
	case common.SubCommandAdd:
		return m.addModel()
	}
	return nil
}

func (m teaModel) listRegisteredModels() func() tea.Msg {
	c := m.edgeClient
	return func() tea.Msg {
		models, err := c.GetModels()
		if err != nil {
			return common.ErrMsg{err}
		}
		return registeredModelsMsg(models)
	}
}

func (m teaModel) addModel() func() tea.Msg {
	c := m.edgeClient
	return func() tea.Msg {
		params, err := pipelines.ReadParams(m.flags[flags.FlagParams.String()])
		if err != nil {
			return common.ErrMsg{err}
		}
		_, err = c.AddNewModelWithImage(
			m.flags[flags.FlagModelName.String()],
			m.flags[flags.FlagModelDescription.String()],
			m.flags[flags.FlagVersionName.String()],
			"",
			params.ToSimpleMap(),
		)
		if err != nil {
			return common.ErrMsg{err}
		}
		return newModelAddedMsg{}

	}
}

// Update updates the model according to the message
func (m teaModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case common.ErrMsg:
		m.err = msg
		return m, tea.Quit

	case registeredModelsMsg:
		m.registeredModelsList = msg
		return m, tea.Quit
	case newModelAddedMsg:
		return m, tea.Quit
	}
	return m, nil
}

// View returns the view corresponding to the subcommand
func (m teaModel) View() string {
	if m.err != nil {
		return common.ErrorStyle.Render(fmt.Sprintf("Error: %s", m.err))
	}
	switch m.subCommand {
	case common.SubCommandList:
		return m.viewListModels()
	case common.SubCommandAdd:
		return common.MessageStyle.Render("\nAdding model information.......") + common.Success.Render("[OK]\n\n")
	}
	return ""
}

func (m teaModel) viewListModels() string {
	columns := []table.Column{
		{Title: "Id", Width: 4},
		{Title: "Name", Width: 20},
		{Title: "Description", Width: 60},
	}

	rows := make([]table.Row, 0)

	if m.registeredModelsList != nil {
		for _, model := range m.registeredModelsList {
			rows = append(
				rows, table.Row{
					model.ID,
					model.Name,
					model.Description,
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
	return common.TableBaseStyle.Render(t.View()) + "\n"
}

// Cmd represents the models command
var Cmd = common.NewCmd(
	"models",
	"Manage models",
	`Manage Open Data Hub models from the command line.

This command will list all the registered models available in the Open Data Hub model registry.`,
	cobra.NoArgs,
	[]flags.Flag{flags.FlagModelRegistryURL.SetParentFlag()},
	common.SubCommandList,
	NewTeaModel,
)

func init() {
	Cmd.AddCommand(addCmd)
}
