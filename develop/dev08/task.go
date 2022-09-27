package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const (
	resetColor = "\033[0m"
	redColor   = "\033[31m"

	cdCommand   = "cd"
	pwdCommand  = "pwd"
	echoCommand = "echo"
	killCommand = "kill"
	psCommand   = "ps"

	quitCommand = "\\quit"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		printPrompt()
		command, args := parseInput(scanner)
		switch command {
		case cdCommand:
			cd(args[0])
		case pwdCommand:
			pwd()
		case echoCommand:
			echo(args)
		case psCommand:
			ps()
		case killCommand:
			kill(args[0])
		case quitCommand:
			return
		default:
			fmt.Printf("%sUnknown command %q, try again\n", redColor, command)
		}
	}
}

func printPrompt() {
	location, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%s%s%s ", resetColor, location, ">")
}

func parseInput(scanner *bufio.Scanner) (string, []string) {
	scanner.Scan()
	text := scanner.Text()
	fields := strings.Fields(text)
	if len(fields) == 0 {
		return "", nil
	}
	if len(fields) == 1 {
		return fields[0], nil
	}

	return fields[0], fields[1:]
}

func cd(to string) {
	if err := os.Chdir(to); err != nil {
		fmt.Printf("%s%s\n", redColor, err)
	}
}

func pwd() {
	location, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("\nPath\n----\n%s\n\n", location)
}

func echo(args []string) {
	var sb strings.Builder

	for i, arg := range args {
		sb.WriteString(arg)
		if i != len(args)-1 {
			sb.WriteRune(' ')
		}
	}

	fmt.Println(sb.String())
}

func kill(process string) {
	cmd := exec.Command("kill", process)
	if runtime.GOOS == "windows" {
		cmd = exec.Command("tskill", process)
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("%s%s\n", redColor, err)
	}
}

func ps() {
	cmd := exec.Command("ls", "-lah")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("tasklist")
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("%s%s\n", redColor, err)
	}
}
