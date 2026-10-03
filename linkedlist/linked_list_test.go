package linkedlist_test

import (
	"golang-snippets/linkedlist"
	"testing"
)

func TestCreateLinkedList(t *testing.T) {
	list, err := linkedlist.CreateLinkedList()
	if err != nil {
		t.Fatalf("CreateLinkedList() error = %v, want nil", err)
	}
	if list == nil {
		t.Fatal("CreateLinkedList() returned a nil list")
	}
}

func ExampleLinkedList_PrintList() {
	list, _ := linkedlist.CreateLinkedList()
	list.PrintList()
	// Output: len: 0
}

func ExampleLinkedList_AddNode() {
	list, _ := linkedlist.CreateLinkedList()
	list.AddNode(1)
	list.AddNode(2)
	list.AddNode(3)
	list.PrintList()
	// Output:
	// len: 3
	// 1->2->3->
}

func ExampleLinkedList_AddNodeDirect() {
	list, _ := linkedlist.CreateLinkedList()
	list.AddNodeDirect(&linkedlist.ListNode{}) // val is unexported, so it stays 0
	list.PrintList()
	// Output:
	// len: 1
	// 0->
}

func TestLinkedListSearchNode(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		search    int
		wantFound bool
		wantPos   int
		wantSteps int
	}{
		{name: "empty list", input: nil, search: 1, wantFound: false, wantPos: 0, wantSteps: 0},
		{name: "single, found", input: []int{5}, search: 5, wantFound: true, wantPos: 0, wantSteps: 0},
		{name: "single, missing", input: []int{5}, search: 3, wantFound: false, wantPos: 1, wantSteps: 1},
		{name: "first", input: []int{2, 1, 3}, search: 2, wantFound: true, wantPos: 0, wantSteps: 0},
		{name: "middle", input: []int{2, 1, 3}, search: 1, wantFound: true, wantPos: 1, wantSteps: 1},
		{name: "last", input: []int{2, 1, 3}, search: 3, wantFound: true, wantPos: 2, wantSteps: 2},
		{name: "missing walks whole list", input: []int{2, 1, 3}, search: 4, wantFound: false, wantPos: 3, wantSteps: 3},
		{name: "duplicates, first match wins", input: []int{7, 4, 7}, search: 7, wantFound: true, wantPos: 0, wantSteps: 0},
		{name: "negative", input: []int{0, -1, 1}, search: -1, wantFound: true, wantPos: 1, wantSteps: 1},
		{name: "insertion order kept", input: []int{9, 1, 5}, search: 5, wantFound: true, wantPos: 2, wantSteps: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list, _ := linkedlist.CreateLinkedList()
			for _, v := range tt.input {
				list.AddNode(v)
			}

			found, pos, steps := list.SearchNode(tt.search)
			if found != tt.wantFound || pos != tt.wantPos || steps != tt.wantSteps {
				t.Errorf("SearchNode(%d) = (%v, %d, %d), want (%v, %d, %d)",
					tt.search, found, pos, steps, tt.wantFound, tt.wantPos, tt.wantSteps)
			}
		})
	}
}
