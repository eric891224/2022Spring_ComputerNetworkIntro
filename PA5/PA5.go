package main

import (
	"fmt"
	"bufio"
	"net"
	"strconv"
	"os"
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
		defer conn.Close()

		reader := bufio.NewReader(conn)
		message, errr := reader.ReadString('\n')
		check(errr)
		fmt.Printf("Upload file size: %s", message)

		outfile, errf := os.Create("whatever.txt")
		check(errf)
		defer outfile.Close()
		line_count := 1
		file_len := 0
		text_len := 0
		size_limit, _ := strconv.Atoi(message[:len(message)-1])
		fwriter := bufio.NewWriter(outfile)
		for {
			text_line, errt := reader.ReadString('\n')
			check(errt)

			str_line_count := strconv.Itoa(line_count)
			_len, errw := fwriter.WriteString(str_line_count + " " + text_line)
			check(errw)

			file_len += _len
			line_count += 1
			text_len += len(text_line)
			if (text_len >= size_limit) {
				break
			}
		}
		fwriter.Flush()
		fmt.Printf("Output file size: %d\n", file_len)

		writer := bufio.NewWriter(conn)
		response := fmt.Sprintf("%d bytes received, %d bytes file generated\n", size_limit, file_len)
		_, err := writer.WriteString(response)
		check(err)
		writer.Flush()
	}
}