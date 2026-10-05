package linkedlist

import "math"

// single node in a linked list
type ListNode struct {
	next *ListNode
	val  int
}

// every list will have a printlist function, add node to list
type ListInterface interface {
	PrintList()
	AddNodeDirect(*ListNode)
	AddNode(val int)

	/*
		3 values in it:
		- found or not - boolean
		- if found, at what index/location/position?
		- how many elements had to be traversed? (this will be equal to location in the start, but after skiplist, this gets drastically reduced)
	*/
	SearchNode(val int) (bool, int, int)

	// is the node deleted or not
	DeleteNode(val int) bool
}

// max neg value allowed
const MAX_NEGATIVE = math.MinInt
