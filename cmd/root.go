package cmd

import (
	"fmt"

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
	err := rootCmd.Execute()
	cobra.CheckErr(err)
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", fmt.Sprintf("config file (default is $HOME/.config/%s/config.json)", appName))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		viper.AddConfigPath(fmt.Sprintf("%s/.config/%s/", home, appName))
		viper.AddConfigPath(fmt.Sprintf("/etc/%s/", appName))
		viper.SetConfigName("config.json")
	}

	viper.SetConfigType("json")
	viper.SetEnvPrefix(appName)
	err := viper.BindEnv("user")
	cobra.CheckErr(err)
	err = viper.BindEnv("password")
	cobra.CheckErr(err)
	err = viper.BindEnv("tenant_id")
	cobra.CheckErr(err)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func run(cmd *cobra.Command, args []string) {
	fmt.Println(appName, "[command]")
}
