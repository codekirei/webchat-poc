package cmd

import (
	"fmt"

	"github.com/codekirei/webchat-poc/backend"
	"github.com/spf13/cobra"
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
	fmt.Println("serve called")
	server.Start()
}
