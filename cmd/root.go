package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tasks",
	Short: "Tasks is a CLI todo app.",
	Long:  "Tasks is a full-featured command line todo app wherein you can add tasks and track them. You can even mark them as completed and set a deadline for each task.",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error occurred while executing Tasks.\n%s\n", err)
		os.Exit(1)
	}
}
