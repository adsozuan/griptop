package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Title struct {
	*tview.TextView
}

func NewTitle(text string) *Title {
	t := Title{
		TextView: tview.NewTextView(),
	}
	t.SetTextColor(tcell.ColorWhite)
	t.SetTextAlign(tview.AlignCenter)
	t.SetBorderPadding(0, 0, 1, 1)
	t.SetText(text)
	t.SetBackgroundColor(tcell.ColorBlue)
	return &t
}

func Run() {
	app := tview.NewApplication()

	grid := tview.NewGrid().SetRows(1, 5, 0).SetColumns(50, 50).
		AddItem(NewTitle("griptop on my PC"),
			0, 0, 1, 2, 0, 0, false).
		AddItem(NewProgressBar("CPU", 43.1),
			1, 0, 1, 1, 0, 0, false).
		AddItem(NewTitle("DYN").
			SetBackgroundColor(tcell.ColorGreen),
			1, 1, 1, 1, 0, 0, false).
		AddItem(NewTitle("STAT").
			SetBackgroundColor(tcell.ColorRed),
			2, 0, 1, 2, 0, 0, false)

	if err := app.SetRoot(grid, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
