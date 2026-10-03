//go:build ignore

// Scratch pad: experiment freely, run with `go run linkedlist/scratch.go`.
// The build tag above keeps it out of the linkedlist package, builds and tests.
package main

import (
	"golang-snippets/linkedlist"
)

func main() {
	list, _ := linkedlist.CreateLinkedList()

	list.AddNode(1)
	list.AddNode(2)
	list.AddNode(3)

	// list.PrintList()
	// fmt.Println()

	list2, _ := linkedlist.CreateSortedList()
	// list2.PrintList()

	testCases := [][]int{
		{2, 1, 3, 0},
		{1, 3, 2},
		{1, 2, 3, 2},
		{5, -1},
	}

	for _, tcase := range testCases {
		list2, _ = linkedlist.CreateSortedList()
		for _, val := range tcase {
			list2.AddNode(val)
		}
		list2.PrintList()

	}

	// fmt.Println()
}
