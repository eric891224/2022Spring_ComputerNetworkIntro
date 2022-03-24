package main

import "fmt"
import "os"
import "bufio"
import "strconv"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	input := ""
	output := ""
	fmt.Printf("Input filename: ")
	fmt.Scanf("%s", &input)
	fmt.Printf("Output filename: ")
	fmt.Scanf("%s", &output)

	inputF, err1 := os.Open(input)
	check(err1)
	defer inputF.Close()

	outputF, err2 := os.Create(output)
	check(err2)
	defer outputF.Close()

	scanner := bufio.NewScanner(inputF)
	writer := bufio.NewWriter(outputF)

	for i:=1; scanner.Scan(); i++ {
		s := strconv.Itoa(i)
		writer.WriteString(s)
		writer.WriteString(" ")
		writer.WriteString(scanner.Text())
		writer.WriteString("\n")
	}
	writer.Flush()
}