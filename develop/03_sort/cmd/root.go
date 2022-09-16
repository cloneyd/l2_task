package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"sort"
	"strings"
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
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Welcome to mysort util!")

			if len(args) == 0 {
				return errors.New("you haven't specified file to sort")
			}

			res, err := readFile(args[0])
			if err != nil {
				return err
			}

			key = 1

			var resCol []string
			for _, line := range res {
				words := strings.Fields(line)[1]
				resCol = append(resCol, words)
			}

			// mathematically
			sort.Sort(MathSort(res))
			fmt.Println("Mathematically:")
			for _, line := range res {
				fmt.Printf("%q\n", line)
			}

			// by length
			sort.Sort(LenSort(res))
			fmt.Println("By length:")
			for _, line := range res {
				fmt.Printf("%q\n", line)
			}

			sort.Sort(sort.Reverse(LenSort(res)))
			fmt.Println("Reverse by length:")
			for _, line := range res {
				fmt.Printf("%q\n", line)
			}

			return nil
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
