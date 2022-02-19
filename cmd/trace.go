package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "trace",
	Short: "IP tracker trace command",
	Long:  `IP tracker trace command. Long description goes here.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running trace command ...")
	},
}
