package main

import "fmt"

func main() {
	fmt.Println("hello world")
	fmt.Println("let's scan a single variable: ----")

	var a int

	fmt.Scan(&a)

	fmt.Printf("scanned a single variable: %v\n", a)
}
