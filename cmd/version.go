package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)
var ver= "1.0.0"
// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Shows the current version of the hackathon-cli",	
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version:",ver)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
