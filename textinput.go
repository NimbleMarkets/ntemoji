package ntemoji

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

// TextInput is a lightweight, zero-dependency text input component for Go and Bubble Tea.
// It avoids importing external clipboard packages to keep compilation lightweight and CG0-free.
type TextInput struct {
	Placeholder string
	Prompt      string

	PromptStyle      lipgloss.Style
	TextStyle        lipgloss.Style
	PlaceholderStyle lipgloss.Style
	CursorStyle      lipgloss.Style

	value   []rune
	cursor  int
	width   int
	focused bool
}

// NewTextInput constructs a TextInput model.
func NewTextInput() TextInput {
	return TextInput{
		width: 15,
	}
}

// SetWidth sets the character width constraint for the text input.
func (t *TextInput) SetWidth(w int) {
	t.width = w
}

// Focus focuses the input.
func (t *TextInput) Focus() tea.Cmd {
	t.focused = true
	return nil
}

// Blur blurs/unfocuses the input.
func (t *TextInput) Blur() {
	t.focused = false
}

// Focused returns true if the input is currently focused.
func (t *TextInput) Focused() bool {
	return t.focused
}

// Value returns the current text value of the input.
func (t *TextInput) Value() string {
	return string(t.value)
}

// SetValue sets the text value of the input.
func (t *TextInput) SetValue(s string) {
	t.value = []rune(s)
	t.cursor = len(t.value)
}

// Update handles keyboard messages to mutate the text input value and cursor.
func (t TextInput) Update(msg tea.Msg) (TextInput, tea.Cmd) {
	if !t.focused {
		return t, nil
	}

	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "backspace":
			if t.cursor > 0 {
				t.value = append(t.value[:t.cursor-1], t.value[t.cursor:]...)
				t.cursor--
			}
		case "delete":
			if t.cursor < len(t.value) {
				t.value = append(t.value[:t.cursor], t.value[t.cursor+1:]...)
			}
		case "left":
			if t.cursor > 0 {
				t.cursor--
			}
		case "right":
			if t.cursor < len(t.value) {
				t.cursor++
			}
		case "home", "ctrl+a":
			t.cursor = 0
		case "end", "ctrl+e":
			t.cursor = len(t.value)
		default:
			// Handle printable character insertions
			text := msg.Key().Text
			if text == "" && msg.Key().Code == ' ' {
				text = " "
			}
			if text != "" {
				runes := []rune(text)
				t.value = append(t.value[:t.cursor], append(runes, t.value[t.cursor:]...)...)
				t.cursor += len(runes)
			}
		}
	}
	return t, nil
}

// View renders the text input string.
func (t TextInput) View() string {
	prompt := t.PromptStyle.Render(t.Prompt)

	var textStr string
	if len(t.value) == 0 && t.Placeholder != "" {
		if t.focused {
			textStr = t.CursorStyle.Render(" ") + t.PlaceholderStyle.Render(t.Placeholder)
		} else {
			textStr = t.PlaceholderStyle.Render(t.Placeholder)
		}
	} else {
		var s strings.Builder
		for i := 0; i < len(t.value); i++ {
			if i == t.cursor && t.focused {
				s.WriteString(t.CursorStyle.Render(string(t.value[i])))
			} else {
				s.WriteRune(t.value[i])
			}
		}
		if t.cursor == len(t.value) && t.focused {
			s.WriteString(t.CursorStyle.Render(" "))
		}
		textStr = t.TextStyle.Render(s.String())
	}

	return prompt + textStr
}
