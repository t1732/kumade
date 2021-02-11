package main

import (
	"os"

	"github.com/cloudfoundry/bytefmt"
	"github.com/olekukonko/tablewriter"
	"github.com/t1732/kumade/cmd"
)

func main() {
	token := cmd.GetToken()
	images := cmd.GetVPCImages(token.ID)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Status", "Size", "CreatedAt"})
	for _, img := range *images {
		table.Append([]string{img.ID, img.Name, img.Status, bytefmt.ByteSize(uint64(img.Size)), img.CreatedAt.String()})
	}
	table.Render()
}
