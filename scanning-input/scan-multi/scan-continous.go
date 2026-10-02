package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func scanAll(a string) []string {
	return strings.Fields(a)
}

func main() {
	/*
		scan - keep reading values constantly and fill them in variables like: Scan(&a, &b, ....)
		scanf - you give the type of the variable too like: Scanf("%s %d", a, b)
		scanln - similar to scan, but only works for the first line: hello/n world only gives hello in the answer

		bufio scanner - creates a buffer pointing
	*/
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter a line: ")
	scanner.Scan()

	scanString := scanner.Text()

	words := scanAll(scanString)

	fmt.Println(words)
	fmt.Println(len(words))
}
