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
