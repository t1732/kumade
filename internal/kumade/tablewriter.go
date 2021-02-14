package kumade

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

type TablewriterWrap struct {
	table *tablewriter.Table
}

func NewWriter() *TablewriterWrap {
	table := tablewriter.NewWriter(os.Stderr)
	table.SetBorder(false)
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t")
	table.SetNoWhiteSpace(true)

	return &TablewriterWrap{table: table}
}

func (t *TablewriterWrap) SetHeader(keys []string) {
	t.table.SetHeader(keys)
}

func (t *TablewriterWrap) Append(data []string) {
	t.table.Append(data)
}

func (t *TablewriterWrap) Render() {
	t.table.Render()
}
