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

func updateUi(app *tview.Application, sysinfoui *SysInfoWidget, prtabledata *ProcessesTableData, sysinfodyn chan services.SystemInfoDyn) {
	for {
		s := <-sysinfodyn

		for _, pr := range s.Processes {
			row := pr.ToString()
			prtabledata.AppendRow(row)
		}

		app.QueueUpdateDraw(func() {
			sysinfoui.cpug.Update(s.CpuUsage)
			sysinfoui.memg.Update(s.MemUsagePercent)
			sysinfoui.tasks.Update(s.TotalTaskCount, s.RunningTaskCount)
			sysinfoui.upt.Update(s.Uptime)
		})
	}
}

func Run(sysinfodyn chan services.SystemInfoDyn, sysinfostatic services.SystemInfoStatic) error {
	app := tview.NewApplication()
	sysinfoui := NewSysInfoWidget(sysinfostatic)
	prtabledata := NewProcessesTableData()
	prtable := tview.NewTable().SetBorders(false).SetSelectable(true, false).SetFixed(1, 1).SetContent(prtabledata)

	grid := tview.NewGrid().SetRows(1, 5, -1).SetColumns(100).
		AddItem(NewTitle("griptop on my PC"),
			0, 0, 1, 2, 0, 0, false).
		AddItem(sysinfoui,
			1, 0, 1, 2, 0, 0, false).
		AddItem(prtable,
			2, 0, 1, 2, 0, 0, true)

	go updateUi(app, sysinfoui, prtabledata, sysinfodyn)
	if err := app.SetRoot(grid, true).EnableMouse(true).Run(); err != nil {
		return err
	}
	return nil
}
