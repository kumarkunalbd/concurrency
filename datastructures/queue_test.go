package datastructres

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueueArray()

	q.Enq(1)
	q.Enq("kunal")
	q.Enq(struct {
		msg   string
		msgid int
	}{
		msgid: 1,
		msg:   "Hello",
	})
	q.Enq(4)
	fmt.Printf("stack=%v\n", q)
	fmt.Printf("Popped=%v\n", q.Deq())
	fmt.Printf("Peek=%v\n", q.Peek())
	q.Enq(9)
	q.Enq(10)
	fmt.Printf("stack=%v\n", q)
	fmt.Printf("Peek=%v\n", q.Peek())
	fmt.Printf("Popped=%v\n", q.Deq())
	fmt.Printf("stack=%v\n", q)
	fmt.Printf("Popped=%v\n", q.Deq())
	fmt.Printf("stack=%v\n", q)
	fmt.Printf("Popped=%v\n", q.Deq())
	fmt.Printf("IsEmpty=%v\n", q.IsEmpty())
	fmt.Printf("stack=%v\n", q)
	fmt.Printf("Popped=%v\n", q.Deq())
	fmt.Printf("stack=%v\n", q)
	fmt.Printf("IsEmpty=%v\n", q.IsEmpty())
	fmt.Printf("Popped=%v\n", q.Deq())
	fmt.Printf("stack=%v\n", q)
	fmt.Printf("IsEmpty=%v\n", q.IsEmpty())
}
