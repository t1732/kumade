package cmd

import (
	"fmt"

	"code.cloudfoundry.org/bytefmt"
	"github.com/spf13/cobra"
	"github.com/t1732/kumade/internal/kumade"
	"github.com/t1732/kumade/pkg/conoha"
)

var (
	imageCmd = &cobra.Command{
		Use:   "image",
		Short: "ConoHa Image API",
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

	command.AddCommand(deleteImageCmd())

	return command
}

func printImages(token string) {
	imgs, err := conoha.Image(token).GetImages()
	cobra.CheckErr(err)

	if len(*imgs) == 0 {
		fmt.Print("no images")
	} else {
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
}

func deleteImageCmd() *cobra.Command {
	var token string

	command := &cobra.Command{
		Use:   "delete [image_id]",
		Short: "ConoHa Image API delete image",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			deleteImage(token, args[0])
		},
	}

	command.PersistentFlags().StringVar(&token, "token", "", "API token")

	return command
}

func deleteImage(token string, imageID string) {
	if err := conoha.Image(token).DeleteImage(imageID); err != nil {
		cobra.CheckErr(err)
	}
}
