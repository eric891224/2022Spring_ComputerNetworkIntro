package main

import "fmt"
import "bufio"
import "net"

func check(e error) {
	if (e != nil) {
		panic(e)
	}
}

func main() {
	fmt.Println("Launching Server...")
	ln, _ := net.Listen("tcp", ":12007") //return socket handle, error
	conn, _ := ln.Accept()	//socket handle ln's class method(accept 1st client request), return socket conn(for data transmission between server and client) and error message
	defer ln.Close()	//socket.Close() to close socket
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	message := ""
	if scanner.Scan() {
		message = scanner.Text()
		fmt.Println(message)
	}

	writer := bufio.NewWriter(conn)
	newline := fmt.Sprintf("%d bytes received\n", len(message))
	_, errw := writer.WriteString(newline)
	check(errw)
	writer.Flush()
}