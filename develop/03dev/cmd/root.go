package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"mysort/internal/file"
	sort2 "mysort/internal/sort"
	"sort"
	"strings"
)

var (
	// Used for flags.
	key         int
	numericSort bool
	reverse     bool
	unique      bool
	check       bool

	noFile     = errors.New("you haven't specified any file")
	invalidKey = errors.New(fmt.Sprintf("Index [%d] out of range", key))

	rootCmd = &cobra.Command{
		Use:   "mysort",
		Short: "Sorting util for strings in FILE(s)",
		Long:  "Write sorted concatenation of all FILE(s) to standard output.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return noFile
			}

			var reader file.Reader

			if unique {
				reader = file.ReadFileUnique
			} else {
				reader = file.ReadFile
			}

			lines, err := file.ReadFiles(args, reader)
			if err != nil {
				return err
			}

			switch {
			case key == 0:
				break
			case key < 0:
				return invalidKey
			default:
				var cols []string
				for _, line := range lines {
					words := strings.Fields(line)
					if len(words) < key {
						return invalidKey
					}
					cols = append(cols, words[key-1])
				}
				lines = cols
			}

			var sorter sort.Interface

			if numericSort {
				sorter = sort2.MathSort(lines)
			} else {
				sorter = sort2.LenSort(lines)
			}

			if reverse {
				sorter = sort.Reverse(sorter)
			}

			sort.Sort(sorter)

			printResult(lines)

			return nil
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func printResult(res []string) {
	for _, line := range res {
		fmt.Println(line)
	}
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&key, "key", "k", 0, "sort via a key")
	rootCmd.PersistentFlags().BoolVarP(&numericSort, "numeric-sort", "n", false, "compare according to string numerical value")
	rootCmd.PersistentFlags().BoolVarP(&reverse, "reverse", "r", false, "reverse the result of comparisons")
	rootCmd.PersistentFlags().BoolVarP(&unique, "unique", "u", false, "output only unique strings")
}
