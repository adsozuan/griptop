package ui

import (
	"github.com/rivo/tview"
)

type SysInfoUi struct {
	grid *tview.Grid
}

func NewSysInfoUi() *SysInfoUi {
	grid := tview.NewGrid()
	grid.SetRows(4, 2)
	//grid.AddItem(NewGauge("CPU", 44.1),
	//	0, 0, 1, 0, 0, 0, false)

	sysInfoUi := SysInfoUi{
		grid: grid,
	}
	return &sysInfoUi
}
