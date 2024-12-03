package datastructres

import (
	"fmt"
	"testing"
)

var minheap = NewMinHeap()

func TestMinHeap_Add(t *testing.T) {

	minheap.Add(&Data{
		Content:  1,
		Priority: 3,
	})

	minheap.Add(&Data{
		Content:  2,
		Priority: 1,
	})

	minheap.Add(&Data{
		Content:  3,
		Priority: 1,
	})

	minheap.Add(&Data{
		Content:  9,
		Priority: 5,
	})

	minheap.Add(&Data{
		Content:  8,
		Priority: 7,
	})

	minheap.Add(&Data{
		Content:  11,
		Priority: 4,
	})

	minheap.Add(&Data{
		Content:  15,
		Priority: 12,
	})

	minheap.Add(&Data{
		Content:  7,
		Priority: 19,
	})

	minheap.Add(&Data{
		Content:  11,
		Priority: 13,
	})

	minheap.Add(&Data{
		Content:  19,
		Priority: 3,
	})

	fmt.Printf("maxharr=%v\n", minheap.DataList)
	fmt.Printf("maxindexes=%v\n", minheap.DataIndexes)
}

func TestMinHeap_RemoveIndexes(t *testing.T) {
	TestMinHeap_Add(t)
	item := minheap.Remove()
	fmt.Printf("1. Removed data =%v Heap=%v IndexMap=%v\n", item, minheap.DataList, minheap.DataIndexes)

	item3 := minheap.RemoveAtIndex(3)
	fmt.Printf("1. Removed data =%v Hedap=%v index=%v\n", item3, minheap.DataList, minheap.DataIndexes)

	item4 := minheap.Remove()
	fmt.Printf("1. Removed data =%v Hedap=%v index=%v\n", item4, minheap.DataList, minheap.DataIndexes)

	item5 := minheap.RemoveAtIndex(5)
	fmt.Printf("1. Removed data =%v Hedap=%v indexes=%v\n", item5, minheap.DataList, minheap.DataIndexes)
}

func TestMinHeap_Remove(t *testing.T) {
	TestMinHeap_Add(t)
	item := minheap.Remove()
	fmt.Printf("1. Removed data =%v Heap=%v IndexMap=%v\n", item, minheap.DataList, minheap.DataIndexes)

	item3 := minheap.Remove()
	fmt.Printf("1. Removed data =%v Hedap=%v index=%v\n", item3, minheap.DataList, minheap.DataIndexes)

	item4 := minheap.Remove()
	fmt.Printf("1. Removed data =%v Hedap=%v index=%v\n", item4, minheap.DataList, minheap.DataIndexes)

	item5 := minheap.Remove()
	fmt.Printf("1. Removed data =%v Hedap=%v indexes=%v\n", item5, minheap.DataList, minheap.DataIndexes)
}
