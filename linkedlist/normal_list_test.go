package linkedlist

import "testing"

// compile-time check that LinkedList satisfies ListInterface
var _ ListInterface = (*LinkedList)(nil)

func newLinkedList(t *testing.T) *LinkedList {
	t.Helper()

	list, err := CreateLinkedList()
	if err != nil {
		t.Fatalf("CreateLinkedList() error = %v, want nil", err)
	}
	ll, ok := list.(*LinkedList)
	if !ok {
		t.Fatalf("CreateLinkedList() returned %T, want *LinkedList", list)
	}
	return ll
}

func TestLinkedListDeleteNode(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		delete int
		wantOK bool
		want   []int
	}{
		{name: "empty list", input: nil, delete: 1, wantOK: false, want: nil},
		{name: "single, found", input: []int{7}, delete: 7, wantOK: true, want: nil},
		{name: "single, missing", input: []int{7}, delete: 5, wantOK: false, want: []int{7}},
		{name: "first", input: []int{3, 4, 5}, delete: 3, wantOK: true, want: []int{4, 5}},
		{name: "middle", input: []int{3, 4, 5}, delete: 4, wantOK: true, want: []int{3, 5}},
		{name: "last", input: []int{3, 4, 5}, delete: 5, wantOK: true, want: []int{3, 4}},
		{name: "missing", input: []int{3, 4, 5}, delete: 9, wantOK: false, want: []int{3, 4, 5}},
		{name: "duplicates, first match wins", input: []int{7, 4, 7}, delete: 7, wantOK: true, want: []int{4, 7}},
		{name: "insertion order kept", input: []int{9, 1, 5}, delete: 1, wantOK: true, want: []int{9, 5}},
		// the dummy head holds MAX_NEGATIVE; delete must never remove it
		{name: "dummy value, not inserted", input: []int{1, 2}, delete: MAX_NEGATIVE, wantOK: false, want: []int{1, 2}},
		{name: "dummy value, inserted for real", input: []int{1, MAX_NEGATIVE}, delete: MAX_NEGATIVE, wantOK: true, want: []int{1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := newLinkedList(t)
			for _, v := range tt.input {
				l.AddNode(v)
			}

			if got := l.DeleteNode(tt.delete); got != tt.wantOK {
				t.Errorf("DeleteNode(%d) = %v, want %v", tt.delete, got, tt.wantOK)
			}
			checkList(t, l, tt.want)
		})
	}
}

// after a delete, appends must still attach to the right tail
func TestLinkedListAddAfterDelete(t *testing.T) {
	l := newLinkedList(t)
	l.AddNode(7)
	l.DeleteNode(7) // empty again
	checkList(t, l, nil)

	l.AddNode(8)
	checkList(t, l, []int{8})

	l.AddNode(9)
	l.DeleteNode(9) // tail removed: tail moves back to 8
	l.AddNode(10)
	checkList(t, l, []int{8, 10})

	l.DeleteNode(8) // head removed
	l.AddNode(11)
	checkList(t, l, []int{10, 11})
}
