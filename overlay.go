package ntemoji

import (
	"strings"

	"github.com/mattn/go-runewidth"
)

// ModalCellSize calculates the cell-width and cell-height of the modal content string.
func ModalCellSize(s string) (width, height int) {
	lines := strings.Split(s, "\n")
	h := len(lines)
	w := 0
	for _, l := range lines {
		lw := runewidth.StringWidth(l)
		if lw > w {
			w = lw
		}
	}
	return w, h
}

// Placement represents the position of the overlay.
type Placement struct {
	top  int
	left int
}

// Fixed returns a Placement configuration.
func Fixed(top, left int) Placement {
	return Placement{top: top, left: left}
}

// ClampedOrigin returns topPad, leftPad centering/clamping the overlay on the background space.
func (p Placement) ClampedOrigin(modalW, modalH, bgW, bgH int) (top, left int) {
	top = p.top
	left = p.left

	if top < 0 {
		top = 0
	}
	if top+modalH > bgH {
		top = bgH - modalH
	}
	if top < 0 {
		top = 0
	}

	if left < 0 {
		left = 0
	}
	if left+modalW > bgW {
		left = bgW - modalW
	}
	if left < 0 {
		left = 0
	}

	return top, left
}

// OverlayView overlays the modalView content on top of the bgView at coordinates (top, left).
func OverlayView(bgView, modalView string, bgW, bgH, top, left int) string {
	bgLines := strings.Split(bgView, "\n")
	fgLines := strings.Split(modalView, "\n")

	// Ensure background is tall enough
	for len(bgLines) < bgH {
		bgLines = append(bgLines, "")
	}

	for i, fgLine := range fgLines {
		y := top + i
		if y < 0 || y >= len(bgLines) {
			continue
		}

		bgLine := bgLines[y]
		curBgW := runewidth.StringWidth(bgLine)
		if curBgW < left {
			bgLine = bgLine + strings.Repeat(" ", left-curBgW)
		}

		// Split the background line around the overlay
		leftPart := runewidth.Truncate(bgLine, left, "")
		rightOffset := left + runewidth.StringWidth(fgLine)
		rightPart := ""
		if curBgW > rightOffset {
			rightPart = runewidth.TruncateLeft(bgLine, rightOffset, "")
		}

		bgLines[y] = leftPart + fgLine + rightPart
	}

	return strings.Join(bgLines, "\n")
}

// CellInModal returns true if the coordinate (x, y) falls inside the modal bounds.
func CellInModal(x, y, top, left, width, height int) bool {
	return x >= left && x < left+width && y >= top && y < top+height
}
