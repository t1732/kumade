package cmd

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/t1732/kumade/internal/identify/tokens"
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
		Short: "Conoha VPC API Access Token",
		Run: func(cmd *cobra.Command, args []string) {
			printToken(short)
		},
	}

	command.PersistentFlags().BoolVarP(&short, "short", "s", false, fmt.Sprintf("Prints %s version info in short format", appName))

	return command
}

func printToken(short bool) {
	response, err := tokens.GetToken()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if short {
		fmt.Print(response.Access.Token.ID)
	} else {
		data := [][]string{
			{response.Access.Token.ID, response.Access.Token.Expires.String()},
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Expired At"})
		table.SetBorder(false)
		table.AppendBulk(data)
		table.Render()
	}
}
