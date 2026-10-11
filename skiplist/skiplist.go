package skiplist

type SkipList struct {
	head  *ListNode
	count int
}

func CreateSkipList() *SkipList {
	// dummy head
	head := &SkipList{}
	return head
}
