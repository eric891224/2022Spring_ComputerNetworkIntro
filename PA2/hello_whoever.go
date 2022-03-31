package main

import "fmt"
import "os"	//include Stdout, Stdin

func main() {
	fmt.Printf("who's there?\n")
	text := ""	//":=" means delcare and assign value at the same time. use "=" to assign value to an existed variable. the data type is automatically determined by the assigned value
	fmt.Scanf("%s", &text)

	fmt.Printf("hello, %s\n", text)
	fmt.Println("hello,", text)
	fmt.Fprintf(os.Stdout, "hello, %s\n", text)
}