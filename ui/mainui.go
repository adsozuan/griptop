package ui

import (
	"adnotanumber.com/griptop/services"
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

func updateUi(app *tview.Application, cpubar *Gauge, sysinfodyn chan services.SystemInfoDyn) {
	for {
		s := <-sysinfodyn
		app.QueueUpdateDraw(func() {
			cpubar.Update(s.CpuUsage)
		})
	}
}

func Run(sysinfodyn chan services.SystemInfoDyn) {
	app := tview.NewApplication()
	cpug := NewGauge("CPU")
	grid := tview.NewGrid().SetRows(1, 5, 0).SetColumns(50, 50).
		AddItem(NewTitle("griptop on my PC"),
			0, 0, 1, 2, 0, 0, false).
		AddItem(cpug,
			1, 0, 1, 1, 0, 0, false).
		AddItem(NewTitle("DYN").
			SetBackgroundColor(tcell.ColorGreen),
			1, 1, 1, 1, 0, 0, false).
		AddItem(NewTitle("STAT").
			SetBackgroundColor(tcell.ColorRed),
			2, 0, 1, 2, 0, 0, false)

	go updateUi(app, cpug, sysinfodyn)
	if err := app.SetRoot(grid, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
