/*
 * Copyright 2024. Open Data Hub Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package models

import (
	"context"
	"fmt"

	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
	openapiclient "github.com/opendatahub-io/ai-edge/cli/pkg/generated/model_registry_client"
	"github.com/opendatahub-io/ai-edge/cli/pkg/tui/keys"
	"github.com/opendatahub-io/ai-edge/cli/pkg/tui/styles"
)

const (
	columnKeyId          = "id"
	columnKeyName        = "name"
	columnKeyDescription = "description"

	minWidth  = 30
	minHeight = 8

	// Add a fixed margin to account for description & instructions
	fixedVerticalMargin = 6
)

type registeredModelsMsg *openapiclient.RegisteredModelList

type errMsg struct{ err error }

func (e errMsg) Error() string { return e.err.Error() }

type Model struct {
	flexTable table.Model
	help      help.Model
	keyMap    *keys.KeyMap
	styles    styles.Styles

	// Window dimensions
	totalWidth  int
	totalHeight int

	// Table dimensions
	horizontalMargin int
	verticalMargin   int

	registeredModelsList *openapiclient.RegisteredModelList
	cursor               int
	choices              int
	err                  error
}

func New(windowSize tea.WindowSizeMsg) tea.Model {
	h := help.New()
	s := styles.DefaultStyles()

	return &Model{
		flexTable: table.New(
			[]table.Column{
				table.NewColumn(columnKeyName, "Name", 10).WithFiltered(true),
				// This table uses flex columns, but it will still need a target
				// width in order to know what width it should fill.  In this example
				// the target width is set below in `recalculateTable`, which sets
				// the table to the width of the screen to demonstrate resizing
				// with flex columns.
				// table.NewFlexColumn(columnKeyElement, "Element", 1),
				table.NewFlexColumn(columnKeyDescription, "Description", 3),
			},
		).
			Filtered(true).
			Focused(true).
			WithBaseStyle(
				s.BaseStyle,
			),
		help:   h,
		keyMap: keys.NewKeyMap(),
		styles: s,

		totalWidth:  windowSize.Width,
		totalHeight: windowSize.Height,
	}

}

func listRegisteredModels() tea.Msg {
	configuration := openapiclient.NewConfiguration()
	configuration.Servers = openapiclient.ServerConfigurations{
		{
			URL: "http://localhost:8080",
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

func (m *Model) Init() tea.Cmd {
	return listRegisteredModels
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	m.flexTable, cmd = m.flexTable.Update(msg)
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			cmds = append(cmds, tea.Quit)
		}
	case tea.WindowSizeMsg:
		m.totalWidth = msg.Width
		m.totalHeight = msg.Height
	case errMsg:
		m.err = msg.err
	case registeredModelsMsg:
		m.registeredModelsList = msg
		m.refreshTable()
	}

	m.recalculateTable()

	return m, tea.Batch(cmds...)
}

func (m *Model) recalculateTable() {
	m.flexTable = m.flexTable.
		WithTargetWidth(m.calculateWidth()).
		WithMinimumHeight(m.calculateHeight())
}

func (m *Model) calculateWidth() int {
	return m.totalWidth - m.horizontalMargin
}

func (m *Model) calculateHeight() int {
	return m.totalHeight - m.verticalMargin - fixedVerticalMargin
}

func (m *Model) View() string {
	title := m.styles.Title.MarginLeft(2).MarginTop(1).Render("Open Data Hub AI Edge CLI / Registered Models")

	help := lipgloss.NewStyle().MarginLeft(4).Render(m.help.View(m.keyMap))

	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		// subTitle,
		m.flexTable.View(),
		help,
	) + "\n"
}

func (m *Model) refreshTable() {
	rows := make([]table.Row, 0)

	for _, model := range m.registeredModelsList.Items {
		rows = append(
			rows, table.NewRow(
				table.RowData{
					columnKeyId:          model.GetId(),
					columnKeyName:        model.GetName(),
					columnKeyDescription: model.GetDescription(),
				},
			),
		)
	}
	m.flexTable = table.New(
		[]table.Column{
			table.NewColumn(columnKeyId, "id", 4),
			table.NewFlexColumn(columnKeyName, "Name", 1).WithFiltered(true),
			table.NewFlexColumn(columnKeyDescription, "Description", 3),
		},
	).
		Filtered(true).
		Focused(true).
		WithRows(rows).
		WithBaseStyle(
			m.styles.BaseStyle,
		)

}
