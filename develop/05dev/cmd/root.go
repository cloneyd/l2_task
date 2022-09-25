package cmd

import (
	"05dev/internal/file"
	"05dev/internal/search"
	"github.com/spf13/cobra"
)

var (
	after      int
	before     int
	context    int
	count      bool
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool

	rootCmd = cobra.Command{
		Use:   "mygrep",
		Short: "Sorting util for strings in FILE(s)",
		Long: `grep searches for PATTERNS in inp. PATTERNS is one or more patterns separated by newline characters,
and grep prints each line that matches a pattern.  Typically PATTERNS should be quoted when grep is used in a
shell command.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			target := args[0]
			lines, err := file.ReadFile(args[1])
			if err != nil {
				return err
			}

			return search.OutputResult(lines, target, after, before, context, count, ignoreCase, invert, fixed, lineNum)
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&after, "after", "A", 0, "Print NUM lines of trailing context after matching lines")
	rootCmd.PersistentFlags().IntVarP(&before, "before", "B", 0, "Print NUM lines of leading context before matching lines.")
	rootCmd.PersistentFlags().IntVarP(&context, "context", "C", 0, "Print NUM lines of output context.")
	rootCmd.PersistentFlags().BoolVarP(&count, "count", "c", false, "Suppress normal output; instead print a count of matching lines for each  input  file.")
	rootCmd.PersistentFlags().BoolVarP(&ignoreCase, "ignore-case", "i", false, "Print NUM lines of trailing context after matching lines")
	rootCmd.PersistentFlags().BoolVarP(&invert, "invert", "v", false, "Print NUM lines of leading context before matching lines.")
	rootCmd.PersistentFlags().BoolVarP(&fixed, "fixed", "F", false, "Print NUM lines of output context.")
	rootCmd.PersistentFlags().BoolVarP(&lineNum, "line-num", "n", false, "Suppress normal output; instead print a count of matching lines for each  input  file.")
}
