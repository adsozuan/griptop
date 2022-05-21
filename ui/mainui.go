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

func updateUi(app *tview.Application, sysinfoui *SysInfoWidget, sysinfodyn chan services.SystemInfoDyn) {
	for {
		s := <-sysinfodyn
		app.QueueUpdateDraw(func() {
			sysinfoui.cpug.Update(s.CpuUsage)
			sysinfoui.memg.Update(s.MemUsagePercent)
			sysinfoui.tasks.Update(s.TotalTaskCount, s.RunningTaskCount)
		})
	}
}

func Run(sysinfodyn chan services.SystemInfoDyn) {
	app := tview.NewApplication()
	sysinfoui := NewSysInfoWidget()
	grid := tview.NewGrid().SetRows(1, 5, -1).
		AddItem(NewTitle("griptop on my PC"),
			0, 0, 1, 2, 0, 0, false).
		AddItem(sysinfoui,
			1, 0, 1, 1, 0, 0, false)

	go updateUi(app, sysinfoui, sysinfodyn)
	if err := app.SetRoot(grid, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
