package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var clusterCmd = &cobra.Command{
	Use:   "watch",
	Short: "watch",
	Long:  "watch for resources and prune them",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
