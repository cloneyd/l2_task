package cmd

import "os"

func openFile(fileName string) (*os.File, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0755)

	if err != nil {
		return nil, err
	}

	return file, nil
}
