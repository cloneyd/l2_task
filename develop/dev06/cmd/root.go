package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var (
	fields    int
	delimiter string
	separated bool

	rootCmd = cobra.Command{
		Use:   "mycut",
		Short: "remove sections from each line of files",
		Long: `Print selected parts of lines from standard output.
       Mandatory arguments to long options are mandatory for short options too.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			lines := readLines()
			splited := getFields(lines)

			return outputFields(splited, fields)
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func readLines() []string {
	var lines []string
	reader := bufio.NewReader(os.Stdin)

	for {
		text, _ := reader.ReadString('\n')

		text = strings.Replace(text, "\n", "", -1)
		text = strings.Replace(text, "\r", "", -1)

		if separated && len(text) == 1 {
			continue
		}

		if len(text) == 0 {
			break
		}

		lines = append(lines, text)
	}

	return lines
}

func getFields(lines []string) [][]string {
	res := make([][]string, len(lines))

	for idx, line := range lines {
		res[idx] = strings.Split(line, delimiter)
	}

	return res
}

func outputFields(splited [][]string, field int) error {
	if field < 1 {
		fmt.Println("invalid field number")
		return errors.New("invalid field number")
	}

	fmt.Println(field)
	for _, line := range splited {
		if field >= len(line) {
			continue
		}

		fmt.Println(line[field-1])
	}

	return nil
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&fields, "fields", "f", 0, "select only these fields; also print any line that contains no delimiter character, unless the -s option is specified")
	rootCmd.PersistentFlags().StringVarP(&delimiter, "delimiter", "d", "\t", "use DELIM instead of TAB for field delimiter")
	rootCmd.PersistentFlags().BoolVarP(&separated, "separated", "s", false, "do not print lines not containing delimiters")
}
