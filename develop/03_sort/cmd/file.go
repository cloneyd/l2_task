package cmd

import (
	"bufio"
	"os"
)

func readFile(fileName string) ([]string, error) {
	var lines []string

	file, err := os.OpenFile(fileName, os.O_RDONLY, 0755)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines, nil
}
