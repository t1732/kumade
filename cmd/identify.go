package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/t1732/kumade/internal/kumade"
	"github.com/t1732/kumade/pkg/conoha"
)

var (
	identifyCmd = &cobra.Command{
		Use:   "identify",
		Short: "identify API",
	}
)

func init() {
	rootCmd.AddCommand(identifyCmd)
	identifyCmd.AddCommand(tokenCmd())
}

func tokenCmd() *cobra.Command {
	var short bool

	command := &cobra.Command{
		Use:   "token",
		Short: "Conoha Identify API create access token",
		Run: func(cmd *cobra.Command, args []string) {
			printToken(short)
		},
	}

	command.PersistentFlags().BoolVarP(&short, "short", "s", false, fmt.Sprintf("Prints %s version info in short format", appName))

	return command
}

func printToken(short bool) {
	response, err := conoha.Identify().CreateToken()
	cobra.CheckErr(err)

	if short {
		fmt.Print(response.Access.Token.ID)
	} else {
		data := [][]string{
			{response.Access.Token.ID, response.Access.Token.Expires.String()},
		}
		table := kumade.NewWriter()
		table.SetHeader([]string{"ID", "Expired At"})
		table.AppendBulk(data)
		table.Render()
	}
}
