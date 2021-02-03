package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	ip := os.Args[1]
	port := os.Args[2]
	conn,err := net.Dial("tcp",fmt.Sprintf( "%s:%s", ip, port))
	defer conn.Close()
	if  err != nil {
		fmt.Print("connect fail\n", err.Error())
	}
	inputReader  := bufio.NewReader(os.Stdin)
	for  {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("read from console failed, err:%v\n", err)
		}
		trimmedInput := strings.TrimSpace(input)
		if trimmedInput == "Q" {
			break
		}
		_, err = conn.Write([]byte(trimmedInput))
		if err != nil {
			fmt.Print("write failed, err: %v\n", err.Error())
		}
	}
}
