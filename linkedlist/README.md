# linkedlist

Step 1 on the way to a Redis-style sorted set: plain lists, so the
O(n) cost of search is visible before skip lists fix it.

## What's here
| File           | What                                                    |
|----------------|---------------------------------------------------------|
| type_defs.go   | ListNode, ListInterface (shared by all lists)           |
| normal_list.go | LinkedList: append-only, real head node                 |
| sorted_list.go | SortedList: sorted insert, dummy head node              |
| scratch.go     | Playground, excluded from builds (`//go:build ignore`)  |

## The dummy head
Both lists start with a dummy node that holds no real data. Real values
start at `head.next`.

```
            head (dummy)          real nodes                    tail
                 │                                                │
                 ▼                                                ▼
            ┌─────────┐      ┌───┐      ┌───┐      ┌───┐      ┌───┐
            │  dummy  │ ───▶ │ 3 │ ───▶ │ 7 │ ───▶ │12 │ ───▶ │19 │ ───▶ nil
            └─────────┘      └───┘      └───┘      └───┘      └───┘
                         pos:  0          1          2          3

empty list:  head ──▶ [ dummy ] ──▶ nil        tail = nil
```

Why: every real node, including the first, has a node before it. So
"delete the first node" works like any other delete: `prev.next = curr.next`,
with no special case for moving `head`.

## SearchNode returns (found, pos, steps)
- pos: 0-based index of the first match; on a miss, the list length
- steps: nodes walked past. Equal to pos for now; the skip list will shrink it

## Run
```
go test ./linkedlist/
go run linkedlist/scratch.go
```

## Next (in this folder)
1. Delete on both lists
2. Resolve head mismatch: LinkedList has a real head, SortedList a dummy one
3. Sorted search: stop early once values pass the target
4. Use `steps` to show search is O(n), the motivation for skip lists

After that, work moves to a new skiplist package.
