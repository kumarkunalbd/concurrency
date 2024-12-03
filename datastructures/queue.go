package datastructres

type Queue interface {
	Enq(interface{})
	Deq() interface{}
	IsEmpty() bool
	Peek() interface{}
	PeekBack() interface{}
}

type QueueArray []interface{}

func NewQueueArray() *QueueArray {
	arr := make(QueueArray, 0)
	return &arr
}

func (q *QueueArray) Enq(obj interface{}) {
	*q = append(*q, obj)
}

func (q *QueueArray) Deq() interface{} {
	item := (*q)[0]
	*q = (*q)[1:]
	return item
}

func (q *QueueArray) Peek() interface{} {
	return (*q)[0]
}

func (q *QueueArray) PeekBack() interface{} {
	return (*q)[len(*q)-1]
}

func (q *QueueArray) IsEmpty() interface{} {
	if len(*q) == 0 {
		return true
	}
	return false
}
