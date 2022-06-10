package ui

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	rowcount   = 200
	columcount = 9
)

var headers [9]string = [9]string{"pid(i)", "user(u)", "pri(r)", "cpu(c)", "mem(m)", "thrd(h)", "disk(d)", "time(t)", "process(p)"}

type ProcessesTableData struct {
	tview.TableContentReadOnly
	data      [rowcount][columcount]string
	currIndex int
}

func NewProcessesTableData() *ProcessesTableData {
	ptd := &ProcessesTableData{}
	ptd.currIndex = 1
	return ptd
}

func (d *ProcessesTableData) GetCell(row int, column int) *tview.TableCell {
	// fill header
	if row == 0 {
		return tview.NewTableCell(fmt.Sprintf("[white::b]%s", strings.ToUpper(headers[column]))).
			SetExpansion(1).
			SetBackgroundColor(tcell.ColorGreen).
			SetTextColor(tcell.ColorWhite).
			SetAlign(tview.AlignLeft).
			SetSelectable(false)
	}

	return tview.NewTableCell(d.data[(row+d.currIndex)%columcount][column])
}

func (d *ProcessesTableData) GetRowCount() int {
	return rowcount
}

func (d *ProcessesTableData) GetColumnCount() int {
	return columcount
}

func (d *ProcessesTableData) AppendRow(row [columcount]string) {
	d.data[d.currIndex] = row
	d.currIndex = (d.currIndex + 1) % columcount
}
