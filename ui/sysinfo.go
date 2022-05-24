package ui

import (
	"adnotanumber.com/griptop/services"
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type SysInfoWidget struct {
	*tview.Grid
	cpug  *Gauge
	memg  *Gauge
	tasks *TasksCount
	ram   *TextWithLabel
	upt   *TextWithLabel
	proc  *TextWithLabel
}

func NewSysInfoWidget(sysinfostatic services.SystemInfoStatic) *SysInfoWidget {
	grid := tview.NewGrid()
	cpug := NewGauge("CPU")
	memg := NewGauge("MEM")
	tasks := NewTaskCount(0, 0)
	ram := NewTextWithLabel("Size:", sysinfostatic.MemSize)
	upt := NewTextWithLabelDyn("Uptime:", "08:34:10")
	proc := NewTextWithLabel("Proc:", sysinfostatic.Proc)

	grid.SetRows(1, 1, 1, 1).SetColumns(-1, -1)
	grid.AddItem(cpug,
		0, 0, 1, 1, 0, 0, false).
		AddItem(memg,
			1, 0, 1, 1, 0, 0, false).
		AddItem(tasks,
			0, 1, 1, 1, 0, 0, false).
		AddItem(ram,
			1, 1, 1, 1, 0, 0, false).
		AddItem(upt,
			2, 1, 1, 1, 0, 0, false).
		AddItem(proc,
			3, 1, 1, 1, 0, 0, false)

	sysInfoUi := SysInfoWidget{
		Grid:  grid,
		cpug:  cpug,
		memg:  memg,
		tasks: tasks,
		ram:   ram,
		upt:   upt,
		proc:  proc,
	}
	return &sysInfoUi
}

type TasksCount struct {
	*TextWithLabel
}

func NewTaskCount(total int, running int) *TasksCount {
	t := TasksCount{
		TextWithLabel: NewTextWithLabel("Tasks:",
			formatTaskCount(total, running)),
	}
	return &t
}

func (t *TasksCount) Update(total int, running int) {
	t.text = formatTaskCount(total, running)
}

func formatTaskCount(total int, running int) string {
	return fmt.Sprintf("%d total, %d runnings", total, running)
}

type TextWithLabel struct {
	*tview.Grid
	label string
	text  string
}

func NewTextWithLabel(label string, text string) *TextWithLabel {
	t := TextWithLabel{
		Grid:  tview.NewGrid(),
		label: label,
		text:  text,
	}

	labelw := tview.NewTextView()
	labelw.SetText(label).SetTextAlign(tview.AlignLeft).SetTextColor(tcell.ColorBlue)

	textw := tview.NewTextView()
	textw.SetText(text).SetTextAlign(tview.AlignLeft).SetTextColor(tcell.ColorWhite)

	t.Grid.SetColumns(-1, -3).
		AddItem(labelw, 0, 0, 1, 1, 0, 0, false).
		AddItem(textw, 0, 1, 1, 1, 0, 0, false)

	return &t
}

func (t *TextWithLabel) Update(text string) {
	t.text = text
}

// NewTextWithLabelDyn Create text widget which must be updated at refresh rate
func NewTextWithLabelDyn(label string, text string) *TextWithLabel {
	t := TextWithLabel{
		Grid:  tview.NewGrid(),
		label: label,
		text:  text,
	}

	t.SetDrawFunc(t.drawTextWithLabel)
	return &t
}

func (t *TextWithLabel) drawTextWithLabel(screen tcell.Screen, x int, y int, width int, height int) (int, int, int, int) {
	centerY := y + height/2
	tview.Print(screen, t.label, x, centerY, width-2, tview.AlignLeft, tcell.ColorBlue)
	tview.Print(screen, t.text, x+9, centerY, width-2, tview.AlignLeft, tcell.ColorWhite)
	return x + 1, centerY + 1, width - 2, height - (centerY + 1 - y)
}
