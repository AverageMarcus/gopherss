package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "gopherss",
		Short: "An RSS reader written in Go",
		RunE: func(cmd *cobra.Command, args []string) error {
			go refreshCmd.RunE(cmd, args)
			return serverCmd.RunE(cmd, args)
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}
