package dev10

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func Telnet() {
	timeout := flag.Duration("timeout", 10*time.Second, "Set telnet timeout")
	flag.Parse()

	if len(os.Args) != 4 {
		fmt.Println("Not enough arguments")
		return
	}
	host := os.Args[2]
	port := os.Args[3]

	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), *timeout)
	if err != nil {
		fmt.Println("Error while connecting:", err)
		return
	}
	defer conn.Close()

	writeDone := make(chan struct{})

	go func() {
		<-writeDone
		io.Copy(os.Stdout, conn)
	}()

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			text := scanner.Text()
			conn.Write([]byte(text + "\n"))
		}
		writeDone <- struct{}{}
	}()

	fmt.Println("Success!")
}
