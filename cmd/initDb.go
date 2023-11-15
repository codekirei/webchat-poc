package cmd

import (
	"log"

	"github.com/codekirei/webchat-poc/backend/constants"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initDbCmd represents the initDb command
var initDbCmd = &cobra.Command{
	Use:   "initDb",
	Short: "Initialize database",
	Run:   initDb,
}

func init() {
	rootCmd.AddCommand(initDbCmd)
}

func initDb(cmd *cobra.Command, args []string) {
	dbDir := viper.GetString(constants.DB_DIR)
	log.Printf("dbdir: %v", dbDir)
}
