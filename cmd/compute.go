package cmd

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/t1732/kumade/internal/conoha"
)

var (
	computeCmd = &cobra.Command{
		Use:   "compute",
		Short: "Compute API",
	}
)

func init() {
	rootCmd.AddCommand(computeCmd)
	computeCmd.AddCommand(serversCmd())
	computeCmd.AddCommand(deleteVMCmd())
}

func serversCmd() *cobra.Command {
	var token string

	command := &cobra.Command{
		Use:   "servers",
		Short: "Compute API server list",
		Run: func(cmd *cobra.Command, args []string) {
			if token == "" {
				token = GetTokenID()
			}
			printServers(token)
		},
	}

	command.PersistentFlags().StringVar(&token, "token", "", "API token")

	return command
}

func printServers(token string) {
	servers, err := conoha.Compute(token).GetServers()
	cobra.CheckErr(err)

	if len(*servers) == 0 {
		fmt.Printf("no servers")
	} else {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Name"})
		table.SetBorder(false)
		for _, sv := range *servers {
			table.Append([]string{
				sv.ID,
				sv.Name,
			})
		}
		table.Render()
	}
}

func deleteVMCmd() *cobra.Command {
	var token string

	command := &cobra.Command{
		Use:   "delete [server_id]",
		Short: "Conoha compute API delete VM",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			deleteVM(token, args[0])
		},
	}

	command.PersistentFlags().StringVar(&token, "token", "", "API token")

	return command
}

func deleteVM(token string, serverID string) {
	if err := conoha.Compute(token).DeleteServer(serverID); err != nil {
		cobra.CheckErr(err)
	}
}
