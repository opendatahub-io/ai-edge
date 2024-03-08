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

package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/opendatahub-io/ai-edge/cli/pkg/tui/home"
	"github.com/opendatahub-io/ai-edge/cli/pkg/tui/keys"
	"github.com/opendatahub-io/ai-edge/cli/pkg/tui/models"
	"github.com/opendatahub-io/ai-edge/cli/pkg/tui/styles"
)

type uiState int

const (
	homeUI uiState = iota
	modelsUI
	imagesUI
	pipelinesUI
	unknown
)

type Model struct {
	models     tea.Model
	home       tea.Model
	keyMap     *keys.KeyMap
	currentUI  uiState
	styles     styles.Styles
	windowSize tea.WindowSizeMsg
}

func NewModel() Model {

	return Model{
		home:      home.New(),
		currentUI: homeUI,
	}
}

func (m *Model) updateKeybindins() {

	switch m.currentUI {
	case homeUI:
		m.keyMap.Enter.SetEnabled(true)
		m.keyMap.Create.SetEnabled(true)
		m.keyMap.Delete.SetEnabled(true)

		m.keyMap.Cancel.SetEnabled(false)

	default:
		m.keyMap.Enter.SetEnabled(true)
		m.keyMap.Create.SetEnabled(true)
		m.keyMap.Delete.SetEnabled(true)
		m.keyMap.Cancel.SetEnabled(false)
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.windowSize = msg
	case home.SelectMessage:
		m.currentUI = modelsUI
		m.models = models.New(m.windowSize)
		cmds = append(cmds, m.models.Init())
	}

	switch m.currentUI {
	case homeUI:
		m.home, cmd = m.home.Update(msg)
	case modelsUI:
		m.models, cmd = m.models.Update(msg)
	}

	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	switch m.currentUI {
	case homeUI:
		return m.home.View()
	case modelsUI:
		return m.models.View()
	default:
		return ""
	}
}
