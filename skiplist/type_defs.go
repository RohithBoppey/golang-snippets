package skiplist

type ListNode struct {
	val  int
	next []*ListNode
}

// type ListInterface interface {
// 	PrintList()
// 	// AddNodeDirect(*ListNode)
// 	// AddNode(val int)

// 	/*
// 		3 values in it:
// 		- found or not - boolean
// 		- if found, at what index/location/position?
// 		- how many elements had to be traversed? (this will be equal to location in the start, but after skiplist, this gets drastically reduced)
// 	*/
// 	SearchNode(val int) (bool, int, int)

// 	// is the node deleted or not
// 	DeleteNode(val int) bool
// }
