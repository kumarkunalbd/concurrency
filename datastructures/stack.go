package datastructres

type Stack interface {
	Push(interface{})
	Pop() interface{}
	IsEmpty() bool
	Peek() interface{}
}

type StackArray []interface{}

func NewStackArray() *StackArray {
	arr := make(StackArray, 0)
	return &arr
}

func (st *StackArray) Push(obj interface{}) {
	*st = append(*st, obj)
}

func (st *StackArray) Peek() interface{} {
	return (*st)[len(*st)-1]
}

func (st *StackArray) Pop() interface{} {
	item := (*st)[len(*st)-1]
	*st = (*st)[:len(*st)-1]
	return item
}

func (st *StackArray) IsEmpty() interface{} {
	if len(*st) == 0 {
		return true
	}
	return false
}
