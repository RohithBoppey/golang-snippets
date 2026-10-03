//go:build ignore

// Scratch pad: experiment freely, run with `go run linkedlist/scratch.go`.
// The build tag above keeps it out of the linkedlist package, builds and tests.
package main

import (
	"fmt"
	"golang-snippets/linkedlist"
)

// all types of linked lists so far are
var llTypes = []func() (linkedlist.ListInterface, error){
	linkedlist.CreateLinkedList,
	linkedlist.CreateSortedList,
}

// some helper functions
func createLinkedListFromArray(inputVals []int) []linkedlist.ListInterface {
	arr := []linkedlist.ListInterface{}
	for _, fn := range llTypes {
		list, _ := fn()
		for _, val := range inputVals {
			list.AddNode(val)
		}
		arr = append(arr, list)
	}
	return arr
}

func searchVal(li linkedlist.ListInterface, val int) {
	found, pos, steps := li.SearchNode(val)
	if found {
		fmt.Printf("Found %v at %v in %v steps\n", val, pos, steps)
	} else {
		fmt.Printf("%v not found\n", val)
	}
}

func main() {
	// list, _ := linkedlist.CreateLinkedList()

	// list.AddNode(1)
	// list.AddNode(2)
	// list.AddNode(3)

	// list.PrintList()
	// fmt.Println()

	// l2, _ := linkedlist.CreateSortedList()
	// l2.PrintList()

	// testCases := [][]int{
	// 	{2, 1, 3, 0},
	// 	{1, 3, 2},
	// 	{1, 2, 3, 2},
	// 	{5, -1},
	// }

	// for _, tcase := range testCases {
	// 	l2, _ = linkedlist.CreateSortedList()
	// 	for _, val := range tcase {
	// 		l2.AddNode(val)
	// 	}
	// 	l2.PrintList()

	// }

	testCases2 := [][]int{
		{2, 1, 3, 0, -1, 9, 11, -2},
	}

	// create both types using the helper function
	allLists := createLinkedListFromArray(testCases2[0])

	for _, list := range allLists {
		list.PrintList()

		searchVal(list, -1)
		searchVal(list, -2)
		searchVal(list, 4)

	}
}
