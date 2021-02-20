package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/t1732/kumade/internal/kumade"
	"github.com/t1732/kumade/pkg/conoha"
)

var (
	networkCmd = &cobra.Command{
		Use:   "network",
		Short: "ConoHa Network API",
	}
)

func init() {
	rootCmd.AddCommand(networkCmd)
	networkCmd.AddCommand(securityGroupsCmd())
	networkCmd.AddCommand(deleteSecurityGroupCmd())
}

func securityGroupsCmd() *cobra.Command {
	var token string
	var infoAll bool

	command := &cobra.Command{
		Use:   "security-groups",
		Short: "ConoHa Network API security group list",
		Run: func(cmd *cobra.Command, args []string) {
			if token == "" {
				token = GetTokenID()
			}
			printSecurityGroups(token, infoAll)
		},
	}

	command.PersistentFlags().StringVar(&token, "token", "", "API token")
	command.PersistentFlags().BoolVarP(&infoAll, "all", "a", false, "Prints security group all info")

	return command
}

func printSecurityGroups(token string, infoAll bool) {
	sgrps, err := conoha.Network(token).GetSecurityGroups()
	cobra.CheckErr(err)

	if len(*sgrps) == 0 {
		fmt.Print("no security groups")
	} else {
		if infoAll {
			table := kumade.NewWriter()
			table.SetHeader([]string{"ID", "Name", "Description"})
			for _, sgrp := range *sgrps {
				table.Append([]string{
					sgrp.ID,
					sgrp.Name,
					sgrp.Description,
				})
			}
			table.Render()
			fmt.Print("\n")

			table = kumade.NewWriter()
			table.SetHeader([]string{"ID", "SecurityGroupId", "Direction", "Ethertype", "PortRangeMax", "PortRangeMin"})
			for _, sgrp := range *sgrps {
				for _, rule := range *sgrp.SecurityGroupRules {
					table.Append([]string{
						rule.ID,
						rule.SecurityGroupId,
						rule.Direction,
						rule.Ethertype,
						fmt.Sprint(rule.PortRangeMax),
						fmt.Sprint(rule.PortRangeMin),
					})
				}
			}
			table.Render()
		} else {
			table := kumade.NewWriter()
			table.SetHeader([]string{"ID", "Name"})
			for _, sgrp := range *sgrps {
				table.Append([]string{
					sgrp.ID,
					sgrp.Name,
				})
			}
			table.Render()
		}
	}
}

func deleteSecurityGroupCmd() *cobra.Command {
	var token string

	command := &cobra.Command{
		Use:   "delete-security-group [security_group_id]",
		Short: "Network API delete security group",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if token == "" {
				token = GetTokenID()
			}
			deleteSecurityGroup(token, args[0])
		},
	}

	command.PersistentFlags().StringVar(&token, "token", "", "API token")

	return command
}

func deleteSecurityGroup(token string, securityGroupID string) {
	if err := conoha.Network(token).DeleteSecurityGroup(securityGroupID); err != nil {
		cobra.CheckErr(err)
	}
}
