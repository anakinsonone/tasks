package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete",
	Args:  cobra.ExactArgs(1),
	Short: "Mark task as complete.",
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("Invalid task ID: %w\n", err)
		}

		if id <= 0 {
			return fmt.Errorf("ID must be a positive integer.\n")
		}

		err = Complete(id)
		if err != nil {
			return fmt.Errorf("Error marking task as complete: %w\n", err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
