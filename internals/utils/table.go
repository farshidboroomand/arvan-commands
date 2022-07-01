package utils

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func NewTable(tableHeaders []string) *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)
	center := 1

	table.SetAlignment(center)
	table.SetHeader(tableHeaders)
	table.SetRowLine(true)
	table.SetRowSeparator("~")

	return table
}
