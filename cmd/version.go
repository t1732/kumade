package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd())
}

func versionCmd() *cobra.Command {
	short := false

	command := &cobra.Command{
		Use:   "version",
		Short: "Print version/build info",
		Run: func(cmd *cobra.Command, args []string) {
			printVersion(short)
		},
	}

	command.PersistentFlags().BoolVarP(&short, "short", "s", false, fmt.Sprintf("Prints %s version info in short format", appName))

	return command
}

func printVersion(short bool) {
	if short {
		fmt.Print(version)
	} else {
		format := "%-10s %s\n"
		fmt.Printf(format, "Version:", version)
		fmt.Printf(format, "Revision:", revision)
	}
}
