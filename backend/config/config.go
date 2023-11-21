package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/codekirei/webchat-poc/backend/constant"
	"github.com/codekirei/webchat-poc/backend/util"
	"github.com/spf13/viper"
)

func setDefaultDbDir() {
	var defaultDbDir string

	xdgDataHome, xdgDataHomeExists := os.LookupEnv("XDG_DATA_HOME")
	if xdgDataHomeExists {
		defaultDbDir = filepath.Join(xdgDataHome, constant.MODULE_NAME)

		// This dir might not exist yet, so let's make sure it does.
		err := util.EnsureDir(defaultDbDir)
		if err != nil {
			log.Fatalf("unable to ensure dir %v: %v", defaultDbDir, err)
		}
	} else {
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatalf("unable to get cwd: %v", err)
		}
		defaultDbDir = cwd
	}

	viper.SetDefault(constant.DB_DIR, defaultDbDir)
}

func Configure() {
	setDefaultDbDir()

	viper.SetConfigName(".config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	xdgConfigHome, xdgConfigHomeExists := os.LookupEnv("XDG_CONFIG_HOME")
	if xdgConfigHomeExists {
		viper.AddConfigPath(
			filepath.Join(xdgConfigHome, constant.MODULE_NAME),
		)
	}

	err := viper.ReadInConfig()
	if err != nil && err != err.(viper.ConfigFileNotFoundError) {
		log.Fatalf("error reading config: %v", err)
	}
}
