package db

import (
	"database/sql"
	"log"
	"path/filepath"

	"github.com/codekirei/webchat-poc/backend/constant"
	"github.com/spf13/viper"
	_ "modernc.org/sqlite"
)

func GetDb() *sql.DB {
	dbDir := viper.GetString(constant.DB_DIR)
	dbFile := filepath.Join(dbDir, "webchat-db")
	db, err := sql.Open("sqlite", dbFile)
	if err != nil {
		log.Fatalf("unable to open sqlite db at %v: %v", dbDir, err)
	}

	return db
}
