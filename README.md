# ntemoji — Emoji picker and swatch keyboard overlay for Bubble Tea

<p>
    <a href="https://github.com/NimbleMarkets/ntemoji/tags"><img src="https://img.shields.io/github/tag/NimbleMarkets/ntemoji.svg" alt="Latest Release"></a>
    <a href="https://pkg.go.dev/github.com/NimbleMarkets/ntemoji?tab=doc"><img src="https://godoc.org/github.com/golang/gddo?status.svg" alt="GoDoc"></a>
    <a href="https://github.com/NimbleMarkets/ntemoji/blob/main/CODE_OF_CONDUCT.md"><img src="https://img.shields.io/badge/Contributor%20Covenant-2.1-4baaaa.svg"  alt="Code Of Conduct"></a>
</p>

`ntemoji` is a [Bubble Tea](https://github.com/charmbracelet/bubbletea) widget that provides a keyboard- and mouse-enabled emoji picker overlay keypad for terminal user interfaces (TUIs). 

It features curated categories of TUI-safe, highly compatible emojis, support for custom palettes, search/filter capabilities, preset shortcuts, and positioning integration via [`bubble-overlay`](https://github.com/madicen/bubble-overlay) and [`bubblezone`](https://github.com/lrstanley/bubblezone).

[Try out the live WASM demo.](https://nimblemarkets.github.io/ntemoji)

## Quickstart

```go
package main

import (
	"fmt"
	"os"

	"github.com/NimbleMarkets/ntemoji"
	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
)

type model struct {
	picker ntemoji.Model
	zm     *zone.Manager
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case ntemoji.EmojiChangedMsg:
		fmt.Printf("Selected emoji: %s\n", msg.Emoji)
		return m, tea.Quit
	case ntemoji.EmojiCanceledMsg:
		return m, tea.Quit
	}

	var cmd tea.Cmd
	updated, cmd := m.picker.Update(msg)
	m.picker = updated.(ntemoji.Model)
	return m, cmd
}

func (m model) View() string {
	return m.zm.Scan(m.picker.View())
}

func main() {
	zm := zone.New()
	picker := ntemoji.New(
		ntemoji.WithInitialEmoji("🚀"),
		ntemoji.WithPresets([]string{"😀", "👍", "🔥", "✅", "❌"}),
		ntemoji.WithShowSearch(true),
	)
	picker.SetZoneManager(zm)

	p := tea.NewProgram(model{picker: picker, zm: zm}, tea.WithAltScreen(), tea.WithMouseAllMotion())
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
```

## SwatchPicker (Popup Modal overlay)

The `SwatchPicker` is a small UI component representing the selected emoji with a dropdown arrow (e.g. `😀 ▼`). Clicking it opens the full emoji keyboard popup centered on the button.

1. **Create the swatch**: `swatch := ntemoji.NewSwatchPicker("🚀", "Trigger")`
2. **Setup bubblezone**: `swatch.SetZoneManager(zm)`
3. **Set bounds & render**: call `swatch.SetBounds(row, col, width, height)` and wrap with overlay rendering:
   ```go
   mainView = swatch.ViewWithOverlay(mainView, width, height)
   return zm.Scan(mainView)
   ```
4. **Update**: forward key and mouse events to the open swatch.

## Keyboard & Mouse Interactions

- **Tab / Shift+Tab**: Cycle focus among the **Presets row**, **Search bar**, **Category selectors**, and the **Emoji grid**.
- **Arrows / hjkl**: Navigate selections within the active presets or the emoji grid. Changing focus to category selection and using ←/→ instantly switches category tabs.
- **Enter**: Confirms the selected emoji and fires an `EmojiChangedMsg`.
- **Escape**: Cancels selection and fires an `EmojiCanceledMsg`.
- **Mouse clicks**: Clicking any category tab switches categories. Hovering over presets/grid items highlights them, and releasing the left-click selects the emoji.

## Configuration Options

Construct the emoji picker with **`New(...Option)`**:

| Option | Purpose |
|--------|---------|
| `WithInitialEmoji(emoji)` | Sets the starting emoji (auto-focuses it in grid/presets). |
| `WithPresets([]string)` | Shows a row of quick-select preset emojis above search/categories. |
| `WithPalettes([]Palette)` | Supplies custom emoji lists and category tabs. |
| `WithStyle(lipgloss.Style)` | Configures custom outer frame border and padding. |
| `WithAutoDismiss(bool)` | When true, selection includes `Dismiss: true` to auto-hide the modal. |
| `WithShowSearch(bool)` | Toggles the search input box for matching keywords. |

## Demos

Built-in demo examples compile out-of-the-box:

```sh
# Run full-screen demo
task build-ex-simple
./bin/simple

# Run multi-swatch grid overlay demo
task build-ex-swatch
./bin/swatch
```

## Vibe coded

This comment was inserted by a human.  Except it wasn't, but then it was reviewed and I wrote this.

## License

[MIT License](./LICENSE.txt) — Copyright (c) 2026 [Neomantra Corp](https://www.neomantra.com).

----
Made with :heart: and :fire: by the team behind [Nimble.Markets](https://nimble.markets).
