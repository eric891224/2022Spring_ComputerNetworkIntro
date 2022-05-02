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

	//write string to conn
	fmt.Fprintf(conn, "HTTP/1.1 404 Not Found\r\n")	//http status line
	fmt.Fprintf(conn, "Date: ...\r\n")				//header line
	fmt.Fprintf(conn, "\r\n")						//empty line indicates end of header
	fmt.Fprintf(conn, "File not found\r\n")			//data line
	fmt.Fprintf(conn, "\r\n")						//end of http response
}