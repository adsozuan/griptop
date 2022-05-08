package ui

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewProgressBar(label string, percentage float64) *tview.Box {
	prgbar := tview.NewTextView().SetDrawFunc(
		func(screen tcell.Screen, x int, y int, width int, height int) (int, int, int, int) {
			value := int(int(percentage) * (width) / 100)
			centerY := y + height/2
			for cx := x + 1; cx < x+value; cx++ {
				screen.SetContent(cx, centerY, tview.BoxDrawingsHeavyVertical,
					nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
			}
			tview.Print(screen, label, x+1, centerY, width-2, tview.AlignLeft, tcell.ColorYellow)
			tview.Print(screen, fmt.Sprintf("%6.2f%%", percentage),
				x+width-10, centerY, width-2, tview.AlignLeft, tcell.ColorYellow)

			return x + 1, centerY + 1, width - 2, height - (centerY + 1 - y)
		})
	return prgbar
}
