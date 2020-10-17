package cmd

import (
	"github.com/averagemarcus/gopherss/internal/feeds"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var refreshCmd = &cobra.Command{
	Use:   "refresh",
	Short: "Refreshes all feeds",
	RunE: func(cmd *cobra.Command, args []string) error {
		return feeds.Refresh()
	},
}

var (
	interval int
)

func init() {
	refreshCmd.Flags().IntVar(&interval, "interval", 15, "How long in minutes to wait before refreshing feeds")
	viper.BindPFlag("refreshTimeoutMinutes", refreshCmd.Flags().Lookup("interval"))

	rootCmd.AddCommand(refreshCmd)
}
