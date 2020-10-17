package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"

	"github.com/averagemarcus/gopherss/cmd"
)

func main() {
	// Load config
	setDefaultConfig()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.gopherss")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Ignore
		} else {
			panic(err)
		}
	} else {
		viper.WatchConfig()
	}

	cmd.Execute()
}

func setDefaultConfig() {
	viper.SetDefault("refreshTimeoutMinutes", 15)
	viper.SetDefault("dbPath", "./feeds.db")
}
