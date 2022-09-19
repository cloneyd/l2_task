package file

import (
	"bufio"
	"os"
)

type Reader func(filepath string) ([]string, error)

func ReadFiles(filepaths []string, reader Reader) ([]string, error) {
	var lines []string

	for _, path := range filepaths {
		fileLines, err := reader(path)
		if err != nil {
			return nil, err
		}
		lines = append(lines, fileLines...)
	}

	return lines, nil
}

func ReadFile(filepath string) ([]string, error) {
	var lines []string

	file, err := os.OpenFile(filepath, os.O_RDONLY, 0755)
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

func ReadFileUnique(filepath string) ([]string, error) {
	unique := make(map[string]struct{})

	file, err := os.OpenFile(filepath, os.O_RDONLY, 0755)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			return nil, err
		}
		if _, ok := unique[line]; !ok {
			unique[line] = struct{}{}
		}
	}

	lines := make([]string, len(unique))

	i := 0
	for key := range unique {
		lines[i] = key
		i++
	}

	return lines, nil
}
