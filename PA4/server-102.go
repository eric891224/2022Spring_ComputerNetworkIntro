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
	ln, _ := net.Listen("tcp", ":12007")
	conn, _ := ln.Accept()
	defer ln.Close()
	defer conn.Close()

	reader := bufio.NewReader(conn)
	message, errr := reader.ReadString('\n')	//must use '' instead of "\n"
	check(errr)
	fmt.Printf("%s", message)

	writer := bufio.NewWriter(conn)
	newline := fmt.Sprintf("%d bytes recerived\n", len(message))
	_, errw := writer.WriteString(newline)
	check(errw)
	writer.Flush()
}