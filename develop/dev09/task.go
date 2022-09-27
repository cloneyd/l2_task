package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"unsafe"
)

func main() {
	if len(os.Args) == 0 {
		fmt.Println("You haven't specified URL")
		return
	}

	url := os.Args[1]

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		file, err := os.Create("index.html")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		body := *(*string)(unsafe.Pointer(&bodyBytes))

		file.WriteString(body)
	}
}
