package main

import (
	"adnotanumber.com/griptop/probes"
	"adnotanumber.com/griptop/ui"
	"fmt"
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

// TODO use https://github.com/shirou/gopsutil

func main() {
	app := tview.NewApplication()
	memUsage := probes.NewMemoryUsage()
	memUsage.Update()
	s := fmt.Sprintf("%v", memUsage)

	grid := tview.NewGrid().SetRows(1, 5, 0).SetColumns(50, 50).
		AddItem(NewTitle("griptop on my PC"),
			0, 0, 1, 2, 0, 0, false).
		AddItem(ui.NewProgressBar("CPU", 43.1),
			1, 0, 1, 1, 0, 0, false).
		AddItem(NewTitle(s).
			SetBackgroundColor(tcell.ColorGreen),
			1, 1, 1, 1, 0, 0, false).
		AddItem(NewTitle("STAT").
			SetBackgroundColor(tcell.ColorRed),
			2, 0, 1, 2, 0, 0, false)

	if err := app.SetRoot(grid, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
