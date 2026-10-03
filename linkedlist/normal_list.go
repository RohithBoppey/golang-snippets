package linkedlist

import "fmt"

// normal linked list
type LinkedList struct {
	head   *ListNode
	tail   *ListNode // to keep track of the ending
	length int
}

func (l *LinkedList) AddNode(val int) {
	// create node and append at the end
	node := &ListNode{next: nil, val: val}
	l.AddNodeDirect(node)
}

func (l *LinkedList) AddNodeDirect(n *ListNode) {
	if l.head == nil {
		// no entries in list
		l.head = n
		l.tail = n
	} else {
		l.tail.next = n
		l.tail = n
	}

	l.length += 1
}

func (l *LinkedList) PrintList() {
	fmt.Printf("len: %d\n", l.length)

	curr := l.head
	for curr != nil {
		fmt.Printf("%d->", curr.val)
		curr = curr.next
	}
}

func CreateLinkedList() (ListInterface, error) {
	// create an empty linked list
	return &LinkedList{nil, nil, 0}, nil
}
