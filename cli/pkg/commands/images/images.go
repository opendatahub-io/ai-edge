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

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	. "github.com/opendatahub-io/ai-edge/cli/pkg/commands/common"
	. "github.com/opendatahub-io/ai-edge/cli/pkg/commands/flags"
	"github.com/opendatahub-io/ai-edge/cli/pkg/edgeclient"
	"github.com/opendatahub-io/ai-edge/cli/pkg/pipelines"
	"github.com/spf13/cobra"
)

type imagesModel struct {
	args        []string
	flags       map[string]string
	pipelineRun edgeclient.PipelineRun
	edgeClient  *edgeclient.Client
	modelImages []edgeclient.ModelImage
	subCommand  SubCommand
	msg         tea.Msg
	err         error
}

func NewImagesModel(
	args []string,
	flags map[string]string,
	subCommand SubCommand,
) *imagesModel {
	return &imagesModel{
		args:       args,
		flags:      flags,
		edgeClient: edgeclient.NewClient(flags[FlagModelRegistryUrl.String()]),
		subCommand: subCommand,
	}
}

func (m imagesModel) listModelImages() func() tea.Msg {
	c := m.edgeClient
	return func() tea.Msg {
		models, err := c.GetModelImages()
		if err != nil {
			return ErrMsg{err}
		}
		return modelImagesMsg(models)
	}
}

func (m imagesModel) syncModelImage() func() tea.Msg {
	c := m.edgeClient
	return func() tea.Msg {
		params, err := pipelines.ReadParams(m.flags[FlagParams.String()])
		if err != nil {
			return ErrMsg{err}
		}
		_, err = c.SyncModelImage(m.args[0], m.args[1], params.ToSimpleMap())
		if err != nil {
			return ErrMsg{err}
		}
		return modelImageSyncedMsg{}
	}

}

func (m imagesModel) buildModelImage() func() tea.Msg {
	c := m.edgeClient
	return func() tea.Msg {

		pipelineRun, err := c.BuildModelImage(
			m.args[0], m.flags[FlagNamespace.String()], m.flags[FlagKubeconfig.String()], nil,
		)
		if err != nil {
			return ErrMsg{err}
		}
		return modelImageBuiltMsg{*pipelineRun}
	}
}

func (m imagesModel) Init() tea.Cmd {
	switch m.subCommand {
	case SubCommandList:
		return m.listModelImages()
	case SubCommandSync:
		return m.syncModelImage()
	case SubCommandBuild:
		return m.buildModelImage()
	}
	return nil
}

func (m imagesModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.msg = msg
	switch msg := msg.(type) {
	case ErrMsg:
		m.err = msg
		return m, tea.Quit

	case modelImagesMsg:
		m.modelImages = msg
		return m, tea.Quit
	case modelImageSyncedMsg:
		return m, tea.Quit
	case modelImageBuiltMsg:
		m.pipelineRun = msg.pipelineRun
		return m, tea.Quit
	}
	return m, nil
}

func (m imagesModel) View() string {
	if m.err != nil {
		return ErrorStyle.Render(fmt.Sprintf("Error: %s", m.err))
	}

	switch m.subCommand {
	case SubCommandList:
		if _, ok := m.msg.(modelImagesMsg); ok {
			return m.viewListModelImages()
		}
	case SubCommandSync:
		if _, ok := m.msg.(modelImageSyncedMsg); ok {
			return MessageStyle.Render("\nModel image synchronized!\n\n")
		}
	case SubCommandBuild:
		if _, ok := m.msg.(modelImageBuiltMsg); ok {
			return lipgloss.JoinVertical(
				lipgloss.Left,
				MessageStyle.Render("\nBuilding model image...")+Success.Render("started\n\n"),
				MessageStyle.Render(
					fmt.Sprintf(
						"Pipeline: %s\tNamespace: %s\n", m.pipelineRun.Name,
						m.pipelineRun.Namespace,
					),
				),
			)
		}
	}
	return ""
}

func (m imagesModel) viewListModelImages() string {
	columns := []table.Column{
		{Title: "Id", Width: 4},
		{Title: "Model Id", Width: 8},
		{Title: "Name", Width: 20},
		{Title: "Description", Width: 40},
		{Title: "Version", Width: 8},
		{Title: "Synced", Width: 6},
		{Title: "URI", Width: 60},
	}

	rows := make([]table.Row, 0)

	if m.modelImages != nil {
		for _, model := range m.modelImages {
			needsSync := CheckMark
			if model.NeedsSync {
				needsSync = WaringSymbol
			}
			rows = append(
				rows, table.Row{
					model.Id,
					model.ModelId,
					model.Name,
					model.Description,
					model.Version,
					needsSync,
					model.URI,
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

var Cmd = NewCmd(
	"images",
	"Manage model container images",
	`Manage Open Data Hub model container images from the command line.

This command allows you to list and build model container images suitable for deployment in edge environments.`,
	cobra.NoArgs,
	[]Flag{FlagNamespace, FlagModelRegistryUrl.SetInherited(), FlagKubeconfig.SetInherited()},
	SubCommandList,
	func(args []string, flags map[string]string, subCommand SubCommand) tea.Model {
		return NewImagesModel(
			args, flags, subCommand,
		)
	},
)

func init() {
	Cmd.PersistentFlags().StringP("namespace", "n", "default", "Namespace to use (e.g. to create a pipeline run)")
	Cmd.AddCommand(syncCmd)
	Cmd.AddCommand(buildCmd)
}
