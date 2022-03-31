package main

import "fmt"
import "os"
import "bufio"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Create("PA2-output.txt")
	check(err)
	defer f.Close()	//defer will make following code be executed to the end of the program (prevent forget to close f)

	writer := bufio.NewWriter(f)
	len, _ := writer.WriteString("This is a test!")	//return length of string ,and error message (_ means ignore)
	fmt.Println(len)
	writer.WriteString("This is a test!")
	writer.Flush()	//put the string from writer to the target file
}