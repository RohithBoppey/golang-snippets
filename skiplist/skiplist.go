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
	fmt.Println("Levels:", h.level)
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

func (h *SkipList) Search(val int) (bool, int) {
	// start from the dummy, so the first real node gets compared too
	currp := h.head

	found := false
	steps := 0

	// outer loop: levels, from the top one down to level 0
	for curr_level := h.level - 1; curr_level >= 0; curr_level -= 1 {
		// inner loop: keep moving right while the next one is still smaller than val
		// when it's nil or >= val, just drop down a level (the outer loop does that)
		for currp.next[curr_level] != nil && currp.next[curr_level].val < val {
			// move to the next one
			currp = currp.next[curr_level]
			steps += 1
		}
	}

	// now currp is the node just before where val would be on level 0
	// so check the next one only once
	if currp.next[0] != nil && currp.next[0].val == val {
		found = true
	}

	return found, steps
}

func CreateSkipList() *SkipList {
	// dummy head, on every level
	head := ListNode{val: 0, next: make([]*ListNode, MAX_LEVEL)}
	sl := &SkipList{head: &head, count: 0, level: 0}
	return sl
}
