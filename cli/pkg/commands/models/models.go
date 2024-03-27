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

package models

import (
	"fmt"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	. "github.com/opendatahub-io/ai-edge/cli/pkg/commands/common"
	. "github.com/opendatahub-io/ai-edge/cli/pkg/commands/flags"
	"github.com/opendatahub-io/ai-edge/cli/pkg/edgeclient"
	"github.com/opendatahub-io/ai-edge/cli/pkg/pipelines"
	"github.com/spf13/cobra"
)

type registeredModelsMsg []edgeclient.Model
type newModelAddedMsg struct{}

type modelsModel struct {
	args                 []string
	flags                map[string]string
	edgeClient           *edgeclient.Client
	registeredModelsList []edgeclient.Model
	err                  error
	subCommand           SubCommand
}

func NewModelsModel(args []string, flags map[string]string, subCommand SubCommand) *modelsModel {
	return &modelsModel{
		args:       args,
		flags:      flags,
		edgeClient: edgeclient.NewClient(flags[FlagModelRegistryUrl.String()]),
		subCommand: subCommand,
	}
}

func (m modelsModel) Init() tea.Cmd {
	switch m.subCommand {
	case SubCommandList:
		return m.listRegisteredModels()
	case SubCommandAdd:
		return m.addModel()
	}
	return nil
}

func (m modelsModel) listRegisteredModels() func() tea.Msg {
	c := m.edgeClient
	return func() tea.Msg {
		models, err := c.GetModels()
		if err != nil {
			return ErrMsg{err}
		}
		return registeredModelsMsg(models)
	}
}

func (m modelsModel) addModel() func() tea.Msg {
	c := m.edgeClient
	return func() tea.Msg {
		params, err := pipelines.ReadParams(m.flags[FlagParams.String()])
		if err != nil {
			return ErrMsg{err}
		}
		_, err = c.AddNewModelWithImage(m.args[0], m.args[1], m.args[2], "", params.ToSimpleMap())
		if err != nil {
			return ErrMsg{err}
		}
		return newModelAddedMsg{}

	}
}

func (m modelsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case ErrMsg:
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

func (m modelsModel) View() string {
	if m.err != nil {
		return ErrorStyle.Render(fmt.Sprintf("Error: %s", m.err))
	}
	switch m.subCommand {
	case SubCommandList:
		return m.viewListModels()
	case SubCommandAdd:
		return MessageStyle.Render("\nModel added successfully\n\n")

	}
	return ""
}

func (m modelsModel) viewListModels() string {
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
					model.Id,
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
	return TableBaseStyle.Render(t.View()) + "\n"
}

// Cmd represents the models command
var Cmd = NewCmd(
	"models",
	"Manage models",
	`Manage Open Data Hub models from the command line.

This command will list all the registered models available in the Open Data Hub model registry.`,
	cobra.NoArgs,
	[]Flag{FlagModelRegistryUrl.SetInherited()},
	SubCommandList,
	func(args []string, flags map[string]string, subCommand SubCommand) tea.Model {
		return NewModelsModel(
			args, flags, subCommand,
		)
	},
)

func init() {
	Cmd.AddCommand(addCmd)
}
