package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var (
	timeout int

	rootCmd = cobra.Command{
		Use: "go-telnet",
		RunE: func(cmd *cobra.Command, args []string) error {
			var sb strings.Builder
			signalChan := make(chan os.Signal, 1)

			signal.Notify(signalChan, syscall.SIGINT)

			if len(args) < 2 {
				return errors.New("you haven't specified host or port")
			}

			host := args[0]
			port := args[1]

			timeoutDuration := time.Duration(timeout) * time.Second
			sb.WriteString(host)
			sb.WriteRune(':')
			sb.WriteString(port)
			address := sb.String()

			conn, err := net.DialTimeout("tcp", address, timeoutDuration)
			if err != nil {
				return err
			}
			defer conn.Close()

			for {
				select {
				case <-signalChan:
					err := conn.Close()
					return err
				default:
					// Чтение входных данных от stdin
					reader := bufio.NewReader(os.Stdin)
					fmt.Print("Text to send: ")
					text, _ := reader.ReadString('\n')
					// Отправляем в socket
					fmt.Fprintf(conn, text+"\n")
					// Прослушиваем ответ
					message, _ := bufio.NewReader(conn).ReadString('\n')
					fmt.Print("Message from server: " + message)
				}
			}
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().IntVar(&timeout, "timeout", 10, "sets awaiting timeout for response")
}
