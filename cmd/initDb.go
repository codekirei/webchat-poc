package cmd

import (
	"log"

	"github.com/codekirei/webchat-poc/backend/db"
	"github.com/spf13/cobra"
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
	_db := db.GetDb()
	_, err := _db.Exec(`
    drop table if exists t;
    create table t(i);
    insert into t values(42), (314);
  `)
	if err != nil {
		log.Fatalf("unable to initialize db: %v", err)
	}
}
