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
	// since dummy is there
	if l.head.next == nil {
		// no entries in list
		l.head.next = n
		l.tail = n
	} else {
		l.tail.next = n
		l.tail = n
	}

	l.length += 1
}

func (l *LinkedList) PrintList() {
	fmt.Printf("len: %d\n", l.length)

	curr := l.head.next
	for curr != nil {
		fmt.Printf("%d->", curr.val)
		curr = curr.next
	}
	fmt.Println()
}

// (found?, pos?, steps?)
func (l *LinkedList) SearchNode(val int) (bool, int, int) {
	// search from left to right and see if the val exists or not
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

// deleted or not?
func (l *LinkedList) DeleteNode(val int) bool {
	// find the first node and delete it
	deleted := false

	if l.head.next == nil {
		// empty list
		return false
	}

	curr, prev := l.head.next, l.head

	if curr.next == nil && curr.val == val {
		// only node
		l.head.next = nil
		l.tail = nil // no nodes left, so no tail either
		l.length -= 1
		return true
	}

	for curr != nil && curr.val != val {
		prev = curr
		curr = curr.next
	}

	if curr == nil {
		// not found
		return false
	}

	// found, so use prev to delete it
	if curr == l.head {
		// real head has no prev, so move head forward instead
		l.head = curr.next
	} else {
		prev.next = curr.next
	}

	if curr == l.tail {
		// deleted the last node, so prev becomes the new tail
		l.tail = prev
	}

	l.length -= 1
	deleted = true

	return deleted
}

func CreateLinkedList() (ListInterface, error) {
	// create an empty linked list
	// always have the dummy node before head
	newLinkedList := &LinkedList{}
	node := &ListNode{next: nil, val: MAX_NEGATIVE}

	newLinkedList.head = node
	newLinkedList.tail = nil

	// create an empty list and return
	return newLinkedList, nil
}
