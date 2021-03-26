package cmd

import (
	m "github.com/allanger/pomoday-backend/third_party/postgres/migrations"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "start http server with configured api",
	Long:  `Starts a http server and serves the configured api`,
	Run: func(cmd *cobra.Command, args []string) {
		m.Migrate()
	},
}

func init() {
	RootCmd.AddCommand(migrateCmd)
	// Default values
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags(	).String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
