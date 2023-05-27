package utils

import (
	"os"
	"github.com/jedib0t/go-pretty/v6/table"
)

func RenderTable(title string, headerNames []string, rows [][]string) string {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	if title != "" {
		t.SetTitle(title)
	}

	hh := []interface{}{}
	for _, header := range headerNames {
		hh = append(hh, header)
	}
	t.AppendHeader(hh)

	for _, row := range rows {
		if len(row) != len(headerNames) {
			continue
		}
		r := []interface{}{}
		for _, rowValue := range row {
			r = append(r, rowValue)
		}
		t.AppendRow(r)
	}
	return t.Render()
}