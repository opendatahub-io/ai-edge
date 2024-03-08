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

package keys

import (
	"github.com/charmbracelet/bubbles/key"
)

type KeyMap struct {
	CursorUp   key.Binding
	CursorDown key.Binding
	Enter      key.Binding
	Create     key.Binding
	Delete     key.Binding
	Cancel     key.Binding
	Quit       key.Binding
	ForceQuit  key.Binding

	State string
}

// TODO: Fix key bindings
func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.CursorUp, k.CursorDown,
	}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.CursorUp, k.CursorDown, k.Enter, k.Create, k.Delete},
	}
}

func NewKeyMap() *KeyMap {
	return &KeyMap{
		CursorUp: key.NewBinding(
			key.WithKeys("ctrl+k"),
			key.WithHelp("ctrl+k", "move up"),
		),
		CursorDown: key.NewBinding(
			key.WithKeys("ctrl+j"),
			key.WithHelp("ctrl+j", "move down"),
		),
		Enter: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "Check out the currently selected branch"),
		),
		Create: key.NewBinding(
			key.WithKeys("ctrl+a"),
			key.WithHelp(
				"ctrl+a",
				"Create a new branch, with confirmation",
			),
		),
		Delete: key.NewBinding(
			key.WithKeys("ctrl+d"),
			key.WithHelp(
				"ctrl+d",
				"Delete the currently selected branch, with confirmation",
			),
		),

		Cancel: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "Cancel"),
		),
	}
}
