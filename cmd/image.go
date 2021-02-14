package cmd

import (
	"os"

	"code.cloudfoundry.org/bytefmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/t1732/kumade/internal/conoha"
)

var (
	imageCmd = &cobra.Command{
		Use:   "image",
		Short: "image API",
	}
)

func init() {
	rootCmd.AddCommand(imageCmd)
	imageCmd.AddCommand(imagesCmd())
}

func imagesCmd() *cobra.Command {
	var token string

	command := &cobra.Command{
		Use:   "images",
		Short: "ConoHa Image API image list",
		Run: func(cmd *cobra.Command, args []string) {
			if token == "" {
				token = GetTokenID()
			}
			printImages(token)
		},
	}

	command.PersistentFlags().StringVar(&token, "token", "", "API token")

	return command
}

func printImages(token string) {
	imgs, err := conoha.Image(token).GetImages()
	cobra.CheckErr(err)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Status", "Visibility", "Size", "Created At", "Updated At"})
	table.SetBorder(false)
	for _, img := range *imgs {
		table.Append([]string{
			img.ID,
			img.Name,
			img.Status,
			img.Visibility,
			bytefmt.ByteSize(uint64(img.Size)),
			img.CreatedAt.String(),
			img.UpdatedAt.String(),
		})
	}
	table.Render()
}
