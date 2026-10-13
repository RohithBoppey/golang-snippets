//go:build ignore

package main

import (
	"fmt"
	"golang-snippets/skiplist"
)

func helperCreateSkiplist(arr []int) *skiplist.SkipList {
	sl := skiplist.CreateSkipList()

	for _, val := range arr {
		sl.AddNode(val)
	}

	return sl
}

func searchVal(li *skiplist.SkipList, val int) {
	found, steps := li.Search(val)
	if found {
		fmt.Printf("Found %v in %v steps\n", val, steps)
	} else {
		fmt.Printf("%v not found\n", val)
	}
}

func main() {
	// [3, 7, 12, 19, 25, 31, 40, 52]
	arr := []int{3, 7, 12, 19, 25, 31, 40, 52}
	sl := helperCreateSkiplist(arr)
	sl.PrintList()

	searchVal(sl, 3)
	searchVal(sl, 25)
	searchVal(sl, 31)
}
