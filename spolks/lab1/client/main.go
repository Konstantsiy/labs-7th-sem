package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const DefaultProtocol = "tcp"

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}

	port := arguments[1]
	c, err := net.Dial(DefaultProtocol, port)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text+"\n")

		message, _ := bufio.NewReader(c).ReadString('\n')
		if message == "exit\n" {
			fmt.Println("TCP client exiting...")
			return
		}
		fmt.Print("->: " + message)
	}
}

