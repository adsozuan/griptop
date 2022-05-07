package ui

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type SysInfoUi struct {
	grid *tview.Grid
}

const (
	prgCell  = "▉"
	prgWidth = 20
	prgWarn  = 13
	prgCrit  = 17
)

func progressUsageString(percentage float64) string {
	progressCell := ""
	value := int(int(percentage) * (prgWidth) / 100)
	for index := 0; index < prgWidth; index++ {
		if index < value {
			progressCell = progressCell + getBarColor(index)

		} else {
			progressCell = progressCell + prgCell
		}
	}
	return progressCell + fmt.Sprintf("%6.2f%%", percentage)
}

func NewSysInfoUi() *SysInfoUi {
	grid := tview.NewGrid()
	grid.SetRows(4, 2)
	grid.AddItem(tview.NewTextView().
		SetText(progressUsageString(44.1)),
		0, 0, 1, 0, 0, 0, false)

	sysInfoUi := SysInfoUi{
		grid: grid,
	}
	return &sysInfoUi
}

func GetColorName(color tcell.Color) string {
	for name, c := range tcell.ColorNames {
		if c == color {
			return name
		}
	}
	return ""
}

func getBarColor(value int) string {

	barCell := ""
	barColor := ""

	if value < prgWarn {
		barColor = GetColorName(tcell.ColorGreen)
	} else if value < prgCrit {
		barColor = GetColorName(tcell.ColorOrange)
	} else {
		barColor = GetColorName(tcell.ColorRed)
	}
	barCell = fmt.Sprintf("[%s::]%s[%s::]", barColor, prgCell, GetColorName(tcell.ColorWhite))
	return barCell
}
