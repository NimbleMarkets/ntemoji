module github.com/NimbleMarkets/ntemoji

go 1.26.3

replace charm.land/bubbletea/v2 => github.com/neomantra/bubbletea/v2 v2.0.0-20260506185856-6506c47fa2f3

// booba-assets stages the WASM demo site (wasm_exec.js + the booba
// terminal runtime) into web/ — see the build-wasm-site Taskfile task.
tool github.com/NimbleMarkets/go-booba/cmd/booba-assets

require (
	charm.land/bubbletea/v2 v2.0.6
	charm.land/lipgloss/v2 v2.0.3
	github.com/lrstanley/bubblezone/v2 v2.0.0
	github.com/mattn/go-runewidth v0.0.23
)

require (
	github.com/NimbleMarkets/go-booba v0.6.1-0.20260519153322-af9fe4ce7d2f // indirect
	github.com/charmbracelet/colorprofile v0.4.3 // indirect
	github.com/charmbracelet/ultraviolet v0.0.0-20260422141423-a0f1f21775f7 // indirect
	github.com/charmbracelet/x/ansi v0.11.7 // indirect
	github.com/charmbracelet/x/term v0.2.2 // indirect
	github.com/charmbracelet/x/termios v0.1.1 // indirect
	github.com/charmbracelet/x/windows v0.2.2 // indirect
	github.com/clipperhouse/displaywidth v0.11.0 // indirect
	github.com/clipperhouse/uax29/v2 v2.7.0 // indirect
	github.com/lucasb-eyer/go-colorful v1.4.0 // indirect
	github.com/muesli/cancelreader v0.2.2 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/spf13/pflag v1.0.10 // indirect
	github.com/xo/terminfo v0.0.0-20220910002029-abceb7e1c41e // indirect
	golang.org/x/sync v0.20.0 // indirect
	golang.org/x/sys v0.43.0 // indirect
)
