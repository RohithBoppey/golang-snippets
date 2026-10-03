package main

import (
	"fmt"
	"os"

	"golang-snippets/scanning"
)

func main() {
	fmt.Print("Enter a line: ")

	scanString, err := scanning.ReadLine(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	words := scanning.ScanWords(scanString)

	fmt.Println(words)
	fmt.Println(len(words))
}
