package main

import (
	"fmt"
	"os"

	"golang-snippets/scanning"
)

func main() {
	fmt.Println("hello world")
	fmt.Println("let's scan a single variable: ----")

	a, err := scanning.ScanInt(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	fmt.Printf("scanned a single variable: %v\n", a)
}
