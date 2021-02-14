package cmd

import (
	"code.cloudfoundry.org/bytefmt"
	"github.com/spf13/cobra"
	"github.com/t1732/kumade/internal/conoha"
	"github.com/t1732/kumade/internal/kumade"
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

	table := kumade.NewWriter()
	table.SetHeader([]string{"ID", "Name", "Status", "Visibility", "Size", "Created At", "Updated At"})
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
