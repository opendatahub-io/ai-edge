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
	"github.com/spf13/cobra"

	"github.com/opendatahub-io/ai-edge/cli/pkg/commands/common"
	"github.com/opendatahub-io/ai-edge/cli/pkg/commands/flags"
	"github.com/opendatahub-io/ai-edge/cli/pkg/edgeclient"
	"github.com/opendatahub-io/ai-edge/cli/pkg/pipelines"
)

type imagesModel struct {
	args          []string
	flags         map[string]string
	pipelineRun   edgeclient.PipelineRun
	edgeClient    *edgeclient.Client
	modelImages   []edgeclient.ModelImage
	subCommand    common.SubCommand
	msg           tea.Msg
	err           error
	selectedImage edgeclient.ModelImage
}

// NewImagesModel creates a new bubbletea model for the images command
func NewImagesModel(
	args []string,
	flgs map[string]string,
	subCommand common.SubCommand,
) tea.Model {
	return &imagesModel{
		args:       args,
		flags:      flgs,
		edgeClient: edgeclient.NewClient(flgs[flags.FlagModelRegistryURL.String()]),
		subCommand: subCommand,
	}
}

func (m imagesModel) listModelImages() func() tea.Msg {
	c := m.edgeClient
	return func() tea.Msg {
		models, err := c.GetModelImages()
		if err != nil {
			return common.ErrMsg{err}
		}
		return modelImagesMsg(models)
	}
}

func (m imagesModel) updateModelImage() func() tea.Msg {
	c := m.edgeClient
	return func() tea.Msg {
		params, err := pipelines.ReadParams(m.flags[flags.FlagParams.String()])
		if err != nil {
			return common.ErrMsg{err}
		}

		_, err = c.UpdateModelImage(
			m.flags[flags.FlagModelID.String()],
			m.flags[flags.FlagVersionName.String()],
			params.ToSimpleMap(),
		)
		if err != nil {
			return common.ErrMsg{err}
		}
		return modelImageSyncedMsg{}
	}

}

func (m imagesModel) buildModelImage() func() tea.Msg {
	c := m.edgeClient
	return func() tea.Msg {
		pipelineRun, err := c.BuildModelImage(
			m.flags[flags.FlagModelID.String()],
			m.flags[flags.FlagVersionName.String()],
			m.flags[flags.FlagNamespace.String()],
			m.flags[flags.FlagKubeconfig.String()],
			nil,
		)
		if err != nil {
			return common.ErrMsg{err}
		}
		return modelImageBuiltMsg{*pipelineRun}
	}
}

func (m imagesModel) describeModelImage() func() tea.Msg {
	c := m.edgeClient
	var modelImage modelImageDescribeMsg
	return func() tea.Msg {
		models, err := c.GetModelImages()
		if err != nil {
			return common.ErrMsg{err}
		}

		for _, model := range models {
			if model.ModelID == m.flags[flags.FlagModelID.String()] && model.Version == m.flags[flags.FlagVersionName.String()] {
				modelImage.selectedImage = model
			}
		}

		return modelImageDescribeMsg(modelImage)
	}
}

func (m imagesModel) Init() tea.Cmd {
	switch m.subCommand {
	case common.SubCommandList:
		return m.listModelImages()
	case common.SubCommandUpdate:
		return m.updateModelImage()
	case common.SubCommandBuild:
		return m.buildModelImage()
	case common.SubCommandDescribe:
		return m.describeModelImage()
	}
	return nil
}

func (m imagesModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.msg = msg
	switch msg := msg.(type) {
	case common.ErrMsg:
		m.err = msg
		return m, tea.Quit
	case modelImagesMsg:
		m.modelImages = msg
		return m, tea.Quit
	case modelImageBuiltMsg:
		m.pipelineRun = msg.pipelineRun
		return m, tea.Quit
	case modelImageSyncedMsg:
		return m, tea.Quit
	case modelImageDescribeMsg:
		m.selectedImage = msg.selectedImage
		return m, tea.Quit
	}
	return m, nil
}

func (m imagesModel) View() string {
	if m.err != nil {
		return common.ErrorStyle.Render(fmt.Sprintf("Error: %s", m.err))
	}

	switch m.subCommand {
	case common.SubCommandList:
		if _, ok := m.msg.(modelImagesMsg); ok {
			return m.viewListModelImages()
		}
	case common.SubCommandUpdate:
		if _, ok := m.msg.(modelImageSyncedMsg); ok {
			return common.MessageStyle.Render("\nUpdating inference container image parameters.......") + common.Success.Render("[OK]\n\n")
		}
	case common.SubCommandBuild:
		if _, ok := m.msg.(modelImageBuiltMsg); ok {
			return lipgloss.JoinVertical(
				lipgloss.Left,
				common.MessageStyle.Render("\nBuilding inference container image.......")+common.Success.Render("[STARTED]\n\n"),
				common.MessageStyle.Render(
					fmt.Sprintf(
						"Pipeline: %s\tNamespace: %s\n", m.pipelineRun.Name,
						m.pipelineRun.Namespace,
					),
				),
			)
		}
	case common.SubCommandDescribe:
		if _, ok := m.msg.(modelImageDescribeMsg); ok {
			return m.viewDescribeModelImages()
		}
	}
	return ""
}

func (m imagesModel) viewListModelImages() string {
	columns := []table.Column{
		{Title: "Model Id", Width: 8},
		{Title: "Name", Width: 20},
		{Title: "Description", Width: 40},
		{Title: "Version", Width: 8},
		{Title: "URI", Width: 60},
	}

	rows := make([]table.Row, 0)

	if m.modelImages != nil {
		for _, model := range m.modelImages {
			rows = append(
				rows, table.Row{
					model.ModelID,
					model.Name,
					model.Description,
					model.Version,
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
	return common.TableBaseStyle.Render(t.View()) + "\n"
}

func (m imagesModel) viewDescribeModelImages() string {

	var parameters []string
	var paramItem string
	for key, value := range m.selectedImage.BuildParams {
		if keySlice, ok := value.([]string); ok {
			for index, v := range keySlice {
				if index == 0 {
					paramItem = common.ParamKeyStyle.Render(fmt.Sprintf("%s:", key)) + fmt.Sprint(v)
					parameters = append(parameters, paramItem)
				} else {
					paramItem = common.ParamKeyStyle.Render("") + fmt.Sprint(v)
					parameters = append(parameters, paramItem)
				}
			}

		} else {
			paramItem = common.ParamKeyStyle.Render(fmt.Sprintf("%s:", key)) + fmt.Sprint(value)
			parameters = append(parameters, paramItem)
		}

	}

	renderView := lipgloss.JoinVertical(
		lipgloss.Left,
		common.TitleStyle.Render("Image Details"),
		common.KeyStyle.Render("Name:")+m.selectedImage.Name,
		common.KeyStyle.Render("Id:")+m.selectedImage.ModelID,
		common.KeyStyle.Render("Description:")+m.selectedImage.Description,
		common.KeyStyle.Render("Version:")+m.selectedImage.Version,
		common.KeyStyle.Render("URI:")+m.selectedImage.URI,
		common.TitleStyle.Render("Parameters:")+"",
	) + fmt.Sprintln("") +
		lipgloss.JoinVertical(
			lipgloss.Left,
			parameters...,
		)
	return renderView + "\n"
}

var Cmd = common.NewCmd(
	"images",
	"Manage inference container images",
	`Manage Open Data Hub inference container images from the command line.

This command allows you to list and build inference container images suitable for deployment in edge environments.`,
	cobra.NoArgs,
	[]flags.Flag{
		flags.FlagNamespace.SetInherited(), flags.FlagModelRegistryURL.SetParentFlag(),
		flags.FlagKubeconfig.SetParentFlag(),
	},
	common.SubCommandList,
	NewImagesModel,
)

func init() {
	Cmd.AddCommand(updateCmd)
	Cmd.AddCommand(buildCmd)
	Cmd.AddCommand(describeCmd)
}
