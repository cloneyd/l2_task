package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	key         int
	numericSort bool
	reverse     bool
	unique      bool

	rootCmd = &cobra.Command{
		Use:   "mysort",
		Short: "Sorting util for strings in FILE(s)",
		Long:  "Write sorted concatenation of all FILE(s) to standard output.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to mysort util!")
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&key, "key", "k", 0, "sort via a key")
	rootCmd.PersistentFlags().BoolVarP(&numericSort, "numeric-sort", "n", false, "compare according to string numerical value")
	rootCmd.PersistentFlags().BoolVarP(&reverse, "reverse", "r", false, "reverse the result of comparisons")
	rootCmd.PersistentFlags().BoolVarP(&unique, "unique", "u", false, "output only unique strings")
}
