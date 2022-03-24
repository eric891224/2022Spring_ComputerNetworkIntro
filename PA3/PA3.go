package main

import "fmt"
import "bufio"
import "os"
import "net"
import "strconv"

func check(e error) {
	if ( e != nil) {
		panic(e)
	}
}

func main() {
	fmt.Println("Connected to Server...")
	conn, errc := net.Dial("tcp", "127.0.0.1:12007")
	check(errc)
	defer conn.Close()
	
	filename := ""
	fmt.Printf("Input filename: ")
	fmt.Scanf("%s", &filename)

	f, errf := os.Open(filename)
	check(errf)
	fi, errs := f.Stat()
	check(errs)
	defer f.Close()

	size := strconv.FormatInt(fi.Size(), 10)
	writer := bufio.NewWriter(conn)
	_, errw := writer.WriteString(size+"\n")
	check(errw)
	fmt.Printf("Send the file size first: %d\n", fi.Size())
	writer.Flush()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		writer.WriteString(scanner.Text()+"\n")
	}
	writer.Flush()

	m_scanner := bufio.NewScanner(conn)
	if m_scanner.Scan() {
		fmt.Printf("Server says: %s\n", m_scanner.Text())
	}
}