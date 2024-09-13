package cmd

import "github.com/spf13/cobra"

var (
	showCompletion bool
	listCmd        = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List all todos.",
		Long:    "List all todos from the database in a tabular format.",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := List(showCompletion)
			if err != nil {
				return err
			}

			return nil
		},
	}
)

func init() {
	listCmd.Flags().BoolVarP(&showCompletion, "all", "a", false, "Show completion status.")
	rootCmd.AddCommand(listCmd)
}
