package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del", "d"},
	Short:   "Delete todo",
	Long:    "Delete a todo with a given ID.",
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("Invalid task ID: %w\n", err)
		}

		if id <= 0 {
			return fmt.Errorf("ID must be a positive integer\n")
		}

		err = Delete(id)
		if err != nil {
			return fmt.Errorf("Error deleting task: %w\n", err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
