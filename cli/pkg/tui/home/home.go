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

package home

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/opendatahub-io/ai-edge/cli/pkg/tui/keys"
	"github.com/opendatahub-io/ai-edge/cli/pkg/tui/styles"
)

const (
	defaultWidth = 30
	listHeight   = 30
)

type SelectMessage struct {
	Selected int
}

type item struct {
	Name        string
	Description string
}

func (i item) FilterValue() string { return i.Name }

type Model struct {
	list   list.Model
	keyMap *keys.KeyMap
}

func New() tea.Model {
	items := []list.Item{
		item{Name: "Models", Description: "View and manage registered models"},
		item{Name: "Model Container Images", Description: "View and manage model container images"},
		item{Name: "Model Pipelines", Description: "View and manage model pipelines"},
	}

	styles := styles.DefaultStyles()
	keys := keys.NewKeyMap()
	l := list.New(items, newItemDelegate(keys, &styles), defaultWidth, listHeight)
	l.Title = "Open Data Hub AI Edge CLI"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.PaginationStyle = styles.Pagination
	l.Styles.HelpStyle = styles.Help

	return &Model{
		list:   l,
		keyMap: keys,
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil
	case tea.KeyMsg:
		switch {

		case key.Matches(msg, m.keyMap.CursorUp):
			m.list.CursorUp()

		case key.Matches(msg, m.keyMap.CursorDown):
			m.list.CursorDown()

		case key.Matches(msg, m.keyMap.Enter):
			cmd = selectCmd(m.list.Cursor())
		}
		cmds = append(cmds, cmd)
	}

	m.list, cmd = m.list.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m *Model) View() string {
	return "\n" + m.list.View()
}

func selectCmd(index int) tea.Cmd {
	return func() tea.Msg {
		return SelectMessage{Selected: index}
	}
}
