package main

import "fmt"
import "bufio"
import "net"
import "os"
// import "io"
import "strconv"

func check(e error) {
	if (e != nil) {
		panic(e)
	}
}

func main() {
	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":12007")
	conn, _ := ln.Accept()
	defer ln.Close()
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	scanner.Scan()
	size := scanner.Text()
	fmt.Printf("Upload file size: %s\n", size)

	outfile, errf := os.Create("whatever.txt")
	check(errf)
	defer outfile.Close()
	line_count := 1
	file_len := 0
	text_len := 0
	size_limit, _ := strconv.Atoi(size)
	fwriter := bufio.NewWriter(outfile)
	for scanner.Scan() {
		str_line_count := strconv.Itoa(line_count)
		_len, errw := fwriter.WriteString(str_line_count + " " + scanner.Text() + "\n")
		check(errw)
		file_len += _len
		line_count += 1
		text_len += len(scanner.Text() + "\n")
		if (text_len >= size_limit) {
			break
		}
	}
	fwriter.Flush()
	fmt.Printf("Output file size: %d\n", file_len)

	writer := bufio.NewWriter(conn)
	response := fmt.Sprintf("%s bytes received, %d bytes file generated\n", size, file_len)
	_, err := writer.WriteString(response)
	check(err)
	writer.Flush()
}