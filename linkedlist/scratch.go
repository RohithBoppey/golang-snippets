//go:build ignore

// Scratch pad: experiment freely, run with `go run linkedlist/scratch.go`.
// The build tag above keeps it out of the linkedlist package, builds and tests.
package main

import (
	"fmt"
	"golang-snippets/linkedlist"
)

func main() {
	list, _ := linkedlist.CreateLinkedList()

	list.AddNode(1)
	list.AddNode(2)
	list.AddNode(3)

	list.PrintList()
	fmt.Println()
}
