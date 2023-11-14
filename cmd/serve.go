package cmd

import (
	"log"

	"github.com/codekirei/webchat-poc/backend/constants"
	"github.com/codekirei/webchat-poc/backend/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start server",
	Run:   serve,
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func serve(cmd *cobra.Command, args []string) {
	log.Print("serve called")
	dbDir := viper.GetString(constants.DB_DIR)
	log.Printf("dbdir: %v", dbDir)
	server.Start()
}
