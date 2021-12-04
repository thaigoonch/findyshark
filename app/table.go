package app

import (
	"os"
	"strconv"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"golang.org/x/term"
)

var (
	t = table.NewWriter()
)

func TableContentResults(result string) {
	resultSlice := strings.Split(result, "\n")
	var fileNames []string
	var lineNums []string
	var results []string
	fileNames, lineNums, results = parseResults(resultSlice)

	removeDuplicateValuesSlice := removeDuplicateValues(fileNames)
	countOfFiles := len(removeDuplicateValuesSlice)
	countOfResults := len(lineNums)
	resultSummary := "found " + strconv.Itoa(countOfResults) + " results in " + strconv.Itoa(countOfFiles) + " files."
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"filename", "ln#", "result"})
	rows := len(fileNames)
	counter := 0
	for r := 0; r < rows; r++ {
		t.AppendSeparator()
		t.AppendRow([]interface{}{fileNames[r], lineNums[r], results[r]})
		counter = counter + 1
	}
	t.AppendFooter(table.Row{resultSummary})
	renderTable()
}

func renderTable() {
	t.SetStyle(table.StyleRounded)
	t.Style().Color.Header = text.Colors{text.BgBlack, text.FgHiGreen}
	t.Style().Color.Footer = text.Colors{text.BgBlack, text.FgHiGreen}

	// don't let the table exceed terminal width
    if term.IsTerminal(0) {
		width, _, err := term.GetSize(0)
		if err == nil {
			t.SetAllowedRowLength(width)
		}
	}
	t.Render()
}
