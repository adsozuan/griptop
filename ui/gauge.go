package ui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Gauge struct {
	*tview.Box
	label string
	value float64
}

func NewGauge(label string) *Gauge {
	return NewGaugeWithDefault(label, 0.0)
}

func NewGaugeWithDefault(label string, value float64) *Gauge {
	gauge := Gauge{
		Box:   tview.NewBox(),
		label: label,
		value: value,
	}
	gauge.SetDrawFunc(gauge.drawGauge)
	return &gauge
}

func (g *Gauge) Update(value float64) {
	g.value = value
}

func (g *Gauge) drawGauge(screen tcell.Screen, x int, y int, width int, height int) (int, int, int, int) {
	value := int(int(g.value) * (width) / 100)
	centerY := y + height/2
	barcol := tcell.ColorGreen
	if g.value > 50 {
		barcol = tcell.ColorOrange
	}
	if g.value > 75 {
		barcol = tcell.ColorRed
	}

	const bar = rune('|')
	for cx := x + 5; cx < x+value; cx++ {
		screen.SetContent(cx, centerY, bar,
			nil, tcell.StyleDefault.Foreground(barcol))
	}
	tview.Print(screen, g.label, x+1, centerY, width-2, tview.AlignLeft, tcell.ColorBlue)
	tview.Print(screen, "[", x+4, centerY, width-2, tview.AlignLeft, tcell.ColorYellow)
	tview.Print(screen, fmt.Sprintf("%6.2f%%", g.value),
		x+width-10, centerY, width-2, tview.AlignLeft, tcell.ColorOlive)
	tview.Print(screen, "]", x+width-3, centerY, width-2, tview.AlignLeft, tcell.ColorYellow)

	return x + 1, centerY + 1, width - 2, height - (centerY + 1 - y)
}
