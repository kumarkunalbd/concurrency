package datastructres

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	st := NewStackArray()

	st.Push(1)
	st.Push("kunal")
	st.Push(struct {
		msg   string
		msgid int
	}{
		msgid: 1,
		msg:   "Hello",
	})
	st.Push(4)
	fmt.Printf("stack=%v\n", st)
	fmt.Printf("Popped=%v\n", st.Pop())
	fmt.Printf("Peek=%v\n", st.Peek())
	st.Push(9)
	st.Push(10)
	fmt.Printf("stack=%v\n", st)
	fmt.Printf("Peek=%v\n", st.Peek())
	fmt.Printf("Popped=%v\n", st.Pop())
	fmt.Printf("stack=%v\n", st)
	fmt.Printf("Popped=%v\n", st.Pop())
	fmt.Printf("stack=%v\n", st)
	fmt.Printf("Popped=%v\n", st.Pop())
	fmt.Printf("IsEmpty=%v\n", st.IsEmpty())
	fmt.Printf("stack=%v\n", st)
	fmt.Printf("Popped=%v\n", st.Pop())
	fmt.Printf("stack=%v\n", st)
	fmt.Printf("IsEmpty=%v\n", st.IsEmpty())
	fmt.Printf("Popped=%v\n", st.Pop())
	fmt.Printf("stack=%v\n", st)
	fmt.Printf("IsEmpty=%v\n", st.IsEmpty())
}
