package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"

	"github.com/averagemarcus/gopherss/cmd"
)

func main() {
	// Load config
	setDefaultConfig()
	viper.AutomaticEnv()

	cmd.Execute()
}

func setDefaultConfig() {
	viper.SetDefault("REFRESH_TIMEOUT", 15)
	viper.SetDefault("DB_PATH", "./feeds.db")
}
