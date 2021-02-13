package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const appName = "kumade"

var (
	version  = "dev"
	revision = "-"
	cfgFile  string
	rootCmd  = &cobra.Command{
		Use:   appName,
		Short: "CLI for ConoHa API",
		Long:  "Kumade is a CLI for running the ConoHa API",
		Run:   run,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", fmt.Sprintf("config file (default is $HOME/.config/%s/config)", appName))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home + fmt.Sprintf(".config/%s/", appName))
		viper.AddConfigPath(fmt.Sprintf("/etc/%s/", appName))
		viper.SetConfigName("config")
	}

	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func run(cmd *cobra.Command, args []string) {
	fmt.Println(appName, "[command]")
}
