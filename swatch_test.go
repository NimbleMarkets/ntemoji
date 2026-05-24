package ntemoji

import (
	"testing"

	tea "charm.land/bubbletea/v2"
)

func TestSwatchPickerInitialState(t *testing.T) {
	s := NewSwatchPicker("🚀", "Launcher")

	if s.Emoji() != "🚀" {
		t.Errorf("Expected initial emoji to be 🚀, got %s", s.Emoji())
	}

	if s.Open() {
		t.Error("Expected SwatchPicker to be closed initially")
	}
}

func TestSwatchPickerBounds(t *testing.T) {
	s := NewSwatchPicker("🚀", "Launcher")
	s.SetBounds(2, 4, 0, 0)

	if s.row != 2 || s.col != 4 {
		t.Errorf("Expected bounds (2, 4), got (%d, %d)", s.row, s.col)
	}

	w, h := s.Size()
	if s.w != w || s.h != h {
		t.Errorf("Expected size (%d, %d), got (%d, %d)", w, h, s.w, s.h)
	}
}

func TestSwatchPickerClickToOpen(t *testing.T) {
	s := NewSwatchPicker("🚀", "Launcher")
	s.SetBounds(0, 0, 4, 1)

	// Simulate left click press inside bounds
	s, _ = s.Update(tea.MouseClickMsg{
		X:      1,
		Y:      0,
		Button: tea.MouseLeft,
	})

	if !s.Open() {
		t.Error("Expected SwatchPicker to be open after click")
	}
}

func TestSwatchPickerEventRoutingWhenOpen(t *testing.T) {
	s := NewSwatchPicker("🚀", "Launcher")
	s.SetBounds(0, 0, 4, 1)

	// Click to open
	s, _ = s.Update(tea.MouseClickMsg{
		X:      1,
		Y:      0,
		Button: tea.MouseLeft,
	})

	// Set picker focus to Grid to test grid cursor routing
	s.picker.Focus = FocusGrid

	// Press right key, should propagate to picker
	origCursorX := s.picker.GridCursorX
	s, _ = s.Update(tea.KeyPressMsg{Code: tea.KeyRight})

	if s.picker.GridCursorX == origCursorX {
		t.Errorf("Expected picker grid cursor to move right, but it remained %d", s.picker.GridCursorX)
	}
}

func TestSwatchPickerConfirmSelection(t *testing.T) {
	s := NewSwatchPicker("🚀", "Launcher")
	s.SetBounds(0, 0, 4, 1)

	// Click to open
	s, _ = s.Update(tea.MouseClickMsg{
		X:      1,
		Y:      0,
		Button: tea.MouseLeft,
	})

	// Confirm with a chosen emoji
	s, cmd := s.Update(EmojiChangedMsg{Emoji: "🔥", Dismiss: true})

	if s.Open() {
		t.Error("Expected SwatchPicker to close on selection confirmation")
	}

	if s.Emoji() != "🔥" {
		t.Errorf("Expected swatch emoji to update to 🔥, got %s", s.Emoji())
	}

	if cmd != nil {
		t.Error("Expected nil command on selection confirmation since it is handled by the swatch")
	}
}
