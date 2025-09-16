package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func Init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	configDir := filepath.Join(homeDir, ".alistcli")

	if err := os.MkdirAll(configDir, 0755); err != nil {
		panic(err)
	}

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(configDir)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			createDefaultConfig()
		} else {
			panic(err)
		}
	}
}

func createDefaultConfig() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	configDir := filepath.Join(homeDir, ".alistcli")
	configFile := filepath.Join(configDir, "config.json")

	viper.Set("alist_pwd", "")

	if err := viper.WriteConfigAs(configFile); err != nil {
		panic(err)
	}
}
