package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/averagemarcus/gopherss/internal/server"
)

var (
	port string
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the web server",
	RunE: func(cmd *cobra.Command, args []string) error {
		return server.Start(port)
	},
}

func init() {
	serverCmd.Flags().StringVarP(&port, "port", "p", "8080", "The port to run the web server on")
	viper.BindPFlag("PORT", refreshCmd.Flags().Lookup("port"))

	rootCmd.AddCommand(serverCmd)
}
