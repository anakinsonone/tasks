package cmd

import (
	"time"

	"github.com/spf13/cobra"
)

var (
	minutes int
	hours   int
	days    int
)

var addCmd = &cobra.Command{
	Use:   "add",
	Args:  cobra.ExactArgs(1),
	Short: "Add todo.",
	Long:  "Add a todo to the database.",
	RunE: func(cmd *cobra.Command, args []string) error {
		task := args[0]
		var dueDate *time.Time

		if minutes > 0 || hours > 0 || days > 0 {
			due := time.Now().
				Add(time.Duration(minutes) * time.Minute).
				Add(time.Duration(hours) * time.Hour).
				Add(time.Duration(days) * 24 * time.Hour)
			dueDate = &due
		}

		err := Add(task, dueDate)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().
		IntVarP(&minutes, "minutes", "m", 0, "Due in specified number of minutes from now.")
	addCmd.Flags().IntVarP(&hours, "hours", "r", 0, "Due in specified number of hours from now.")
	addCmd.Flags().IntVarP(&days, "days", "d", 0, "Due in specified number of days from now.")
}
