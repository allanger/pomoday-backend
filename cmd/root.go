package cmd

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// HelpCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use: "help",
	// TODO: write decription
	Long: ``,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	serveCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config.yaml)")
	viper.SetDefault("app", "pomoday-backend")

	//Database default creds
	viper.SetDefault("database_username", "user")
	viper.SetDefault("database_password", "12345")
	viper.SetDefault("database_name", "pomoday")
	viper.SetDefault("database_host", "localhost")
	viper.SetDefault("database_port", "5432")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Config file
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("config.yaml")
	}

	// Env vars
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		viper.WatchConfig()
		viper.OnConfigChange(func(e fsnotify.Event) {
			fmt.Println("updated")
		})

	}

}
