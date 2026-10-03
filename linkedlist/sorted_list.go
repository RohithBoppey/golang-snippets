package linkedlist

import "fmt"

// implementing sorted list - always insert the values in the sorted order while printing

type SortedList struct {
	// overriding the methods of a normal list
	LinkedList
}

func (l *SortedList) AddNode(val int) {
	l.AddNodeDirect(&ListNode{val: val})
}

func (l *SortedList) AddNodeDirect(n *ListNode) {
	prev, curr := l.head, l.head.next

	if curr == nil {
		// first element
		prev.next = n
		l.tail = n
	} else {
		// keep iterating the linked list until you find the right spot
		for curr.next != nil && curr.val < n.val {
			prev = curr
			curr = curr.next
		}

		if curr == l.tail && curr.val < n.val {
			// last index
			curr.next = n
			l.tail = n
		} else {
			// in between element
			prev.next = n
			n.next = curr
		}
	}

	l.length += 1
}

func (l *SortedList) PrintList() {
	// fmt.Printf("len: %d\n", l.length)

	curr := l.head.next
	for curr != nil {
		fmt.Printf("%d->", curr.val)
		curr = curr.next
	}
	fmt.Println("")
}

// actual sorted list search + skipList implementation comes in here
// but for now, it's gonna be same as parent, but from head.next step
func (l *SortedList) SearchNode(val int) (bool, int, int) {
	curr := l.head.next
	pos := 0
	found := false
	for curr != nil {
		if curr.val == val {
			found = true
			break
		}
		curr = curr.next
		pos += 1
	}
	return found, pos, pos
}

// constructor
func CreateSortedList() (ListInterface, error) {
	// always have the dummy node before head
	newLinkedList := &SortedList{}
	node := &ListNode{next: nil, val: MAX_NEGATIVE}

	newLinkedList.head = node

	// create an empty list and return
	return newLinkedList, nil
}
