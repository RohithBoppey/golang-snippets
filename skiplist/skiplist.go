package skiplist

import "fmt"

// max levels any node can have (Redis uses 32)
const MAX_LEVEL = 32

type SkipList struct {
	head  *ListNode
	count int // count of total number of nodes in this
	level int // levels in use right now: the tallest real node decides it
}

func printHelper(l *ListNode, level int) {
	// print just one particular level
	for l != nil {
		fmt.Printf("%v->", l.val)
		l = l.next[level]
	}
}

func (h *SkipList) PrintList() {
	// print level wise in this, only the levels in use
	for curr := 0; curr < h.level; curr += 1 {
		fmt.Printf("L%v: ", curr)
		printHelper(h.head.next[curr], curr) // skip the dummy, start from first real node
		fmt.Println("")
	}
}

// temporary, hand-built levels: append only (input must already be sorted)
// every node on level 0, every 2nd also on level 1, every 4th also on level 2
// random levels replace this in step 4
func (h *SkipList) AddNode(val int) {
	levels := 1
	if h.count%2 == 0 {
		levels = 2
	}
	if h.count%4 == 0 {
		levels = 3
	}

	// create the node, with one slot per level it's on
	node := &ListNode{val: val, next: make([]*ListNode, levels)}

	// add to the end of every level this node is on
	for lvl := 0; lvl < levels; lvl += 1 {
		curr := h.head

		for curr.next[lvl] != nil {
			curr = curr.next[lvl]
		}

		curr.next[lvl] = node
	}

	if levels > h.level {
		h.level = levels
	}
	h.count += 1
}

func CreateSkipList() *SkipList {
	// dummy head, on every level
	head := ListNode{val: 0, next: make([]*ListNode, MAX_LEVEL)}
	sl := &SkipList{head: &head, count: 0, level: 0}
	return sl
}
