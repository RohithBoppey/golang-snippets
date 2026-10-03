package linkedlist

import (
	"math"
	"slices"
	"testing"
)

// compile-time check that SortedList satisfies ListInterface
var _ ListInterface = (*SortedList)(nil)

// values walks the list (skipping the dummy head) and returns its values.
// It stops after length+1 nodes so a broken link can't loop forever.
func values(t *testing.T, l *SortedList) []int {
	t.Helper()

	var got []int
	for curr := l.head.next; curr != nil; curr = curr.next {
		got = append(got, curr.val)
		if len(got) > l.length+1 {
			t.Fatalf("list has more nodes than length %d, possible cycle: %v...", l.length, got)
		}
	}
	return got
}

// checkList verifies the values, length and tail of l.
func checkList(t *testing.T, l *SortedList, want []int) {
	t.Helper()

	if got := values(t, l); !slices.Equal(got, want) {
		t.Errorf("values = %v, want %v", got, want)
	}
	if l.length != len(want) {
		t.Errorf("length = %d, want %d", l.length, len(want))
	}

	if len(want) == 0 {
		if l.tail != nil {
			t.Errorf("tail = %v, want nil for empty list", l.tail)
		}
		return
	}
	if l.tail == nil {
		t.Fatalf("tail = nil, want node with val %d", want[len(want)-1])
	}
	if l.tail.val != want[len(want)-1] {
		t.Errorf("tail.val = %d, want %d", l.tail.val, want[len(want)-1])
	}
	if l.tail.next != nil {
		t.Errorf("tail.next = %v, want nil", l.tail.next)
	}
}

func newSortedList(t *testing.T) *SortedList {
	t.Helper()

	list, err := CreateSortedList()
	if err != nil {
		t.Fatalf("CreateSortedList() error = %v, want nil", err)
	}
	sl, ok := list.(*SortedList)
	if !ok {
		t.Fatalf("CreateSortedList() returned %T, want *SortedList", list)
	}
	return sl
}

func TestCreateSortedList(t *testing.T) {
	l := newSortedList(t)

	if l.head == nil {
		t.Fatal("head = nil, want dummy node")
	}
	if l.head.next != nil {
		t.Errorf("head.next = %v, want nil", l.head.next)
	}
	checkList(t, l, nil)
}

func TestSortedListAddNode(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{name: "single", input: []int{5}, want: []int{5}},
		{name: "ascending", input: []int{1, 2, 3}, want: []int{1, 2, 3}},
		{name: "descending", input: []int{3, 2, 1}, want: []int{1, 2, 3}},
		{name: "insert at front", input: []int{1, 2, 3, 0}, want: []int{0, 1, 2, 3}},
		{name: "insert before tail", input: []int{1, 3, 2}, want: []int{1, 2, 3}},
		{name: "insert in middle", input: []int{1, 5, 9, 4}, want: []int{1, 4, 5, 9}},
		{name: "duplicates", input: []int{1, 2, 3, 2}, want: []int{1, 2, 2, 3}},
		{name: "all same", input: []int{7, 7, 7}, want: []int{7, 7, 7}},
		{name: "negative after positive", input: []int{5, -1}, want: []int{-1, 5}},
		{name: "negatives only", input: []int{-3, -1, -2}, want: []int{-3, -2, -1}},
		{name: "zero", input: []int{0, -1, 1}, want: []int{-1, 0, 1}},
		{
			name:  "mixed",
			input: []int{11, 1, 2, -1, 2, 4, 3, 0, -2, 5, 9},
			want:  []int{-2, -1, 0, 1, 2, 2, 3, 4, 5, 9, 11},
		},
		{
			name:  "int extremes",
			input: []int{0, math.MaxInt, math.MinInt},
			want:  []int{math.MinInt, 0, math.MaxInt},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := newSortedList(t)
			for _, v := range tt.input {
				l.AddNode(v)
			}
			checkList(t, l, tt.want)
		})
	}
}

func TestSortedListAddNodeDirect(t *testing.T) {
	l := newSortedList(t)

	l.AddNodeDirect(&ListNode{val: 4})
	l.AddNodeDirect(&ListNode{val: 2})
	l.AddNodeDirect(&ListNode{val: 8})

	checkList(t, l, []int{2, 4, 8})
}

func TestSortedListTailAfterAppend(t *testing.T) {
	// tail must keep up when values keep going after the current tail
	l := newSortedList(t)

	l.AddNode(1)
	l.AddNode(3)
	checkList(t, l, []int{1, 3})

	l.AddNode(2) // before the tail: tail stays 3
	checkList(t, l, []int{1, 2, 3})

	l.AddNode(10) // after the tail: tail becomes 10
	checkList(t, l, []int{1, 2, 3, 10})
}

func ExampleSortedList_PrintList() {
	list, _ := CreateSortedList()
	for _, v := range []int{3, -1, 2} {
		list.AddNode(v)
	}
	list.PrintList()
	// Output: -1->2->3->
}
