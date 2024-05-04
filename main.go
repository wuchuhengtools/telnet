package main

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
)

func parseCliArgs() (string, string) {
	// Create a new CLI app.
	app := &cli.App{
		Name:      "telnet",
		Usage:     "A simple telnet client",
		UsageText: "telnet <address> <port>",
		Action: func(ctx *cli.Context) error {
			// Get the port and check if it is valid.
			port, err := strconv.ParseInt(ctx.Args().Get(1), 10, 64)
			if port < 1 || port > 65535 || err != nil {
				return fmt.Errorf("invalid port")
			}
			// Get the address and check if it is valid.
			address := ctx.Args().First()
			// Check the address must be like: 0.0.0.0 format.
			if net.ParseIP(address) == nil {
				return fmt.Errorf("invalid address")
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

	ip := os.Args[1]
	port := os.Args[2]
	return ip, port
}

func main() {

	// Get the IP and port from the command line arguments
	ip, port := parseCliArgs()
	listener, err := net.Dial("tcp", fmt.Sprintf("%s:%s", ip, port))
	defer listener.Close()
	if err != nil {
		fmt.Print("connect fail\n", err.Error())
		return
	}
	closeApp := sync.WaitGroup{}
	closeApp.Add(1)
	// Receive the coming data
	go receiveComingData(listener, closeApp)

	// Send the input data to the connection.
	go sendInput(listener, closeApp)
	closeApp.Wait()
}

// Receive the coming data
func receiveComingData(listener net.Conn, closeApp sync.WaitGroup) {
	defer func() {
		listener.Close()
		closeApp.Done()
	}()
	reader := bufio.NewReader(listener)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Print("read failed, err: %v\n", err.Error())
			return
		}
		fmt.Print(msg)
	}
}

// Send the input data to the connection.
func sendInput(listener net.Conn, closeApp sync.WaitGroup) {
	// Read from the console
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("read from console failed, err:%v\n", err)
		}
		trimmedInput := strings.TrimSpace(input)
		// If the input is Ctrl+c, then exit the program.
		if trimmedInput == "\x03" {
			break
		}
		_, err = listener.Write([]byte(input))
		if err != nil {
			fmt.Print("write failed, err: %v\n", err.Error())
			closeApp.Done()
		}
	}
}
