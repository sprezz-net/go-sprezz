package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sprezz-net/go-sprezz/internal"
)

func init() {
	rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:   "server",
	Short: "Start server application",
	Run: func(cmd *cobra.Command, args []string) {
		internal.RunServer(cmd.Context())
	},
}
