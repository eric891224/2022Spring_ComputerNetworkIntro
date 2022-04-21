package main

import (
	"fmt"
	"bufio"
	"os"
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

	for {
		conn, _ := ln.Accept()

		reader := bufio.NewReader(conn)
		req, err := http.ReadRequest(reader)
		check(err)

		filename := req.RequestURI[1:len(req.RequestURI)]
		f, errf := os.Open(filename)
		defer f.Close()
		if errf != nil {
			fmt.Println("File not found")
		} else {
			fi, errs := f.Stat()
			check(errs)

			fmt.Printf("File size = %d\n", fi.Size())
		}
		
		conn.Close()
	}
}