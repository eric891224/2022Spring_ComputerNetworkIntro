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
	f, err := os.Open("hello_world.go")
	check(err)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	f.Close()
}