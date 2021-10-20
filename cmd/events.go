package cmd

import (
	"github.com/spf13/cobra"
	"github.com/Dhairya3124/hackathon-cli/hackathonslist"
)

// eventsCmd represents the events command
var eventsCmd = &cobra.Command{
	Use:   "events",
	Short: "Shows the events of the current month at your CLI",
	Run: func(cmd *cobra.Command, args []string) {
		hackathonslist.ListofEvents()
	},
}

func init() {
	rootCmd.AddCommand(eventsCmd)
}
