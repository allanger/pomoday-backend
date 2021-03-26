package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/allanger/pomoday-backend/api"
	m "github.com/allanger/pomoday-backend/third_party/postgres/migrations"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start user-admin service",
	Long: `
	Dresup-admin-users-service
	Description //TODO
	`,
	Run: func(cmd *cobra.Command, args []string) {
		m.Migrate()
		api.Serve()
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
	viper.SetDefault("log_level", "info")
	viper.SetDefault("port", ":8080")
}
