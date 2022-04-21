package main

import (
	"fmt"
	"bufio"
	"net"
	"net/http"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":12007")
	defer ln.Close()
	conn, _ := ln.Accept()
	defer conn.Close()

	reader := bufio.NewReader(conn)
	req, err := http.ReadRequest(reader)
	check(err)

	fmt.Printf("Method: %s\n", req.Method)
	fmt.Printf("Host: %s\n", req.Host)
	fmt.Printf("User-Agent: %s\n", req.UserAgent())
	// fmt.Printf("Request File: %s\n", req.RequestURI)
}