package stack

import "sync"

type Stack interface {
	Push(item Item)
	Pop() *Item
	IsEmpty() bool
	Size() int
	GetFirstItem() *Item
}

type Item interface {}

var _ Stack = &sliceStack{}

type sliceStack struct {
	items []Item
	lock sync.RWMutex
}

func NewSliceStack() Stack {
	return &sliceStack{
		items: []Item{}, // slice的零值是nil，所以需要初始化一个空slice
	}
}

func (ss *sliceStack) Push(item Item) {
	ss.lock.Lock()
	ss.items = append(ss.items, item)
	ss.lock.Unlock()
}

func (ss *sliceStack) Pop() *Item {
	ss.lock.Lock()
	popItem := ss.items[len(ss.items)-1]
	ss.items = ss.items[:len(ss.items)-1]
	ss.lock.Unlock()
	return &popItem
}

func (ss *sliceStack) IsEmpty() bool {
	ss.lock.Lock()
	isEmpty := len(ss.items) == 0
	ss.lock.Unlock()
	return isEmpty
}

func (ss *sliceStack) Size() int {
	ss.lock.Lock()
	size := len(ss.items)
	ss.lock.Unlock()
	return size
}

func (ss *sliceStack) GetFirstItem() *Item {
	ss.lock.Lock()
	popItem := ss.items[len(ss.items)-1]
	ss.lock.Unlock()
	return &popItem
}

