package linkedlist

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
}
