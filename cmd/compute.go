package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/t1732/kumade/internal/kumade"
	"github.com/t1732/kumade/pkg/conoha"
)

var (
	computeCmd = &cobra.Command{
		Use:   "compute",
		Short: "ConoHa Compute API",
	}
)

func init() {
	rootCmd.AddCommand(computeCmd)
	computeCmd.AddCommand(flavorsCmd())
	computeCmd.AddCommand(serversCmd())
	computeCmd.AddCommand(serversDetailCmd())
}

func flavorsCmd() *cobra.Command {
	var token string

	command := &cobra.Command{
		Use:   "flavors",
		Short: "Compute API flavor list",
		Run: func(cmd *cobra.Command, args []string) {
			if token == "" {
				token = GetTokenID()
			}
			printFlavors(token)
		},
	}

	command.PersistentFlags().StringVar(&token, "token", "", "API token")

	return command
}

func printFlavors(token string) {
	flavors, err := conoha.Compute(token).GetFlavors()
	cobra.CheckErr(err)

	if len(*flavors) == 0 {
		fmt.Printf("no flavors")
	} else {
		table := kumade.NewWriter()
		table.SetHeader([]string{"ID", "Name"})
		for _, sv := range *flavors {
			table.Append([]string{
				sv.ID,
				sv.Name,
			})
		}
		table.Render()
	}
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

	command.AddCommand(deleteServerCmd())
	command.AddCommand(serversDetailCmd())

	return command
}

func printServers(token string) {
	servers, err := conoha.Compute(token).GetServers()
	cobra.CheckErr(err)

	if len(*servers) == 0 {
		fmt.Printf("no servers")
	} else {
		table := kumade.NewWriter()
		table.SetHeader([]string{"ID", "Name"})
		for _, sv := range *servers {
			table.Append([]string{sv.ID, sv.Name})
		}
		table.Render()
	}
}

func serversDetailCmd() *cobra.Command {
	var token string

	command := &cobra.Command{
		Use:   "detail",
		Short: "Compute API server detail list",
		Run: func(cmd *cobra.Command, args []string) {
			if token == "" {
				token = GetTokenID()
			}
			printServersDetail(token)
		},
	}

	command.PersistentFlags().StringVar(&token, "token", "", "API token")

	return command
}

func printServersDetail(token string) {
	servers, err := conoha.Compute(token).GetServersDetial()
	cobra.CheckErr(err)

	if len(*servers) == 0 {
		fmt.Printf("no servers")
	} else {
		table := kumade.NewWriter()
		table.SetHeader([]string{"ID", "Name", "Status", "MacAddr", "IP", "SecurityGroups"})
		for _, sv := range *servers {
			for _, add := range sv.Addresses {
				table.Append([]string{
					sv.ID,
					sv.Name,
					sv.Status,
					add.OsExtIPSMacAddr,
					add.IP,
				})
			}
		}
		table.Render()
	}
}

func MapSecurityGroups(securityGroups *[]conoha.SecurityGroup) {
	names := []*string{}
	for _, e := range *securityGroups {
		names = append(names, e.Name)
	}
	return strings.Join(names, ",")
}

func deleteServerCmd() *cobra.Command {
	var token string

	command := &cobra.Command{
		Use:   "delete [server_id]",
		Short: "Compute API delete server",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			deleteServer(token, args[0])
		},
	}

	command.PersistentFlags().StringVar(&token, "token", "", "API token")

	return command
}

func deleteServer(token string, serverID string) {
	if err := conoha.Compute(token).DeleteServer(serverID); err != nil {
		cobra.CheckErr(err)
	}
}
