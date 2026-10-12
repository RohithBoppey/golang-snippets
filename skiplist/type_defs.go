package skiplist

// each node contains address of it's next pointer but for every level
/*
[3, 7, 12, 19, 25, 31, 40, 52]
level 0: dummy -> [3, 7, 12, 19, 25, 31, 40, 52]
level 1: dummy -> [3, 12, 25, 40]
level 2: dummy -> [3, 25]

the dummy head is on every level, real nodes only on their own levels:
dummy - next = [->3, ->3, ->3]
3     - next = [->7, ->12, ->25]
7     - next = [->12]
*/

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
