package main

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

func main() {
	app := tview.NewApplication()
	layout := tview.NewFlex().AddItem(NewTitle("griptop"), 0, 1, false).
		AddItem(tview.NewFlex().AddItem(tview.NewBox(), 0, 5, false),
			0, 5, false)

	if err := app.SetRoot(layout, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
