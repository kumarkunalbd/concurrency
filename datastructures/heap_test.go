package datastructres

import (
	"fmt"
	"testing"
)

var mxheap = NewMaxHeap()

func TestMaxHeap_Add(t *testing.T) {

	mxheap.Add(&Data{
		Content:  1,
		Priority: 3,
	})

	mxheap.Add(&Data{
		Content:  2,
		Priority: 1,
	})

	mxheap.Add(&Data{
		Content:  3,
		Priority: 1,
	})

	mxheap.Add(&Data{
		Content:  9,
		Priority: 5,
	})

	mxheap.Add(&Data{
		Content:  8,
		Priority: 7,
	})

	mxheap.Add(&Data{
		Content:  11,
		Priority: 4,
	})

	mxheap.Add(&Data{
		Content:  15,
		Priority: 12,
	})

	mxheap.Add(&Data{
		Content:  7,
		Priority: 19,
	})

	mxheap.Add(&Data{
		Content:  11,
		Priority: 13,
	})

	mxheap.Add(&Data{
		Content:  19,
		Priority: 3,
	})

	fmt.Printf("maxharr=%v\n", mxheap.DataList)
	fmt.Printf("maxindexes=%v\n", mxheap.DataIndexes)
}

func TestMaxHeap_RemoveIndexes(t *testing.T) {
	TestMaxHeap_Add(t)
	item := mxheap.Remove()
	fmt.Printf("1. Removed data =%v Heap=%v IndexMap=%v\n", item, mxheap.DataList, mxheap.DataIndexes)

	item3 := mxheap.RemoveAtIndex(3)
	fmt.Printf("1. Removed data =%v Hedap=%v index=%v\n", item3, mxheap.DataList, mxheap.DataIndexes)

	item4 := mxheap.Remove()
	fmt.Printf("1. Removed data =%v Hedap=%v index=%v\n", item4, mxheap.DataList, mxheap.DataIndexes)

	item5 := mxheap.RemoveAtIndex(5)
	fmt.Printf("1. Removed data =%v Hedap=%v indexes=%v\n", item5, mxheap.DataList, mxheap.DataIndexes)
}

func TestMaxHeap_Remove(t *testing.T) {
	TestMaxHeap_Add(t)
	item := mxheap.Remove()
	fmt.Printf("1. Removed data =%v Heap=%v IndexMap=%v\n", item, mxheap.DataList, mxheap.DataIndexes)

	item3 := mxheap.Remove()
	fmt.Printf("1. Removed data =%v Hedap=%v index=%v\n", item3, mxheap.DataList, mxheap.DataIndexes)

	item4 := mxheap.Remove()
	fmt.Printf("1. Removed data =%v Hedap=%v index=%v\n", item4, mxheap.DataList, mxheap.DataIndexes)

	item5 := mxheap.Remove()
	fmt.Printf("1. Removed data =%v Hedap=%v indexes=%v\n", item5, mxheap.DataList, mxheap.DataIndexes)
}
