package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const VERSION string = "0.0.1"

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Kumade",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(VERSION)
	},
}
