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

package styles

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	Title lipgloss.Style

	NormalTitle lipgloss.Style
	NormalDesc  lipgloss.Style

	SelectedTitle lipgloss.Style
	SelectedDesc  lipgloss.Style

	BaseStyle lipgloss.Style

	Pagination lipgloss.Style
	Help       lipgloss.Style
	QuitText   lipgloss.Style
}

func DefaultStyles() (s Styles) {
	return PurpleStyles()
}

func PurpleStyles() (s Styles) {
	s.Title = lipgloss.NewStyle().
		Background(lipgloss.Color("62")).
		Foreground(lipgloss.Color("230")).
		Padding(0, 1)

	s.NormalTitle = lipgloss.NewStyle().
		PaddingLeft(4).
		Foreground(lipgloss.AdaptiveColor{Light: "#1A1A1A", Dark: "#DDDDDD"})

	s.NormalDesc = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#A49FA5", Dark: "#777777"})

	s.SelectedTitle = lipgloss.NewStyle().
		PaddingLeft(2).
		Foreground(lipgloss.AdaptiveColor{Light: "#EE6FF8", Dark: "#EE6FF8"})

	s.SelectedDesc = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#F793FF", Dark: "#AD58B4"})

	s.BaseStyle = lipgloss.NewStyle().
		BorderForeground(lipgloss.Color("#a38")).
		Foreground(lipgloss.AdaptiveColor{Light: "#1A1A1A", Dark: "#DDDDDD"}).
		Align(lipgloss.Left)

	s.Pagination = list.DefaultStyles().
		PaginationStyle.
		PaddingLeft(4)

	s.Help = list.DefaultStyles().
		HelpStyle.
		PaddingLeft(4).
		PaddingBottom(1)

	s.QuitText = lipgloss.NewStyle().
		Margin(1, 0, 2, 4)

	return s
}

func GreenStyles() (s Styles) {
	s.Title = lipgloss.NewStyle().
		Background(lipgloss.Color("62")).
		Foreground(lipgloss.Color("230")).
		Padding(0, 1)

	s.NormalTitle = lipgloss.NewStyle().
		PaddingLeft(4).
		Foreground(lipgloss.AdaptiveColor{Light: "#1A1A1A", Dark: "#DDDDDD"})

	s.NormalDesc = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#A49FA5", Dark: "#777777"})

	s.SelectedTitle = lipgloss.NewStyle().
		PaddingLeft(2).
		Foreground(lipgloss.AdaptiveColor{Light: "#EE6FF8", Dark: "#EE6FF8"})

	s.SelectedDesc = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#F793FF", Dark: "#AD58B4"})

	s.BaseStyle = lipgloss.NewStyle().
		BorderForeground(lipgloss.Color("#04B575")).
		// Foreground(lipgloss.Color("#a7a")).
		Align(lipgloss.Left)

	s.Pagination = list.DefaultStyles().
		PaginationStyle.
		PaddingLeft(4)

	s.Help = list.DefaultStyles().
		HelpStyle.
		PaddingLeft(4).
		PaddingBottom(1)

	s.QuitText = lipgloss.NewStyle().
		Margin(1, 0, 2, 4)

	return s
}
