//go:build ignore

package main

import "golang-snippets/skiplist"

func helperCreateSkiplist(arr []int) *skiplist.SkipList {
	sl := skiplist.CreateSkipList()

	for _, val := range arr {
		sl.AddNode(val)
	}

	return sl
}

func main() {
	// [3, 7, 12, 19, 25, 31, 40, 52]
	arr := []int{3, 7, 12, 19, 25, 31, 40, 52}
	sl := helperCreateSkiplist(arr)
	sl.PrintList()
}
