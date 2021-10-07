package queue

import "sync"

type Queue interface {
	Push(item Item)
	Pop() *Item
	IsEmpty() bool
	Size() int
	GetFirstItem() *Item
}

// 定义队列元素
type Item interface{}

// 声明的切片 Queue 实现了 Queue interface
var _ Queue = &sliceQueue{}

type sliceQueue struct {
	items []Item
	lock  sync.RWMutex // 队列读写锁，确保线程安全
}

// New一个切片队列，当然队列还有其他的实现方式，可以使用switch方式（类似工厂模式），New不同的队列
func NewSliceQueue() Queue {
	return &sliceQueue{
		items:  []Item{},
	}
}

func (sq *sliceQueue) Push(item Item) {
	sq.lock.Lock()
	sq.items = append(sq.items, item)
	sq.lock.Unlock()
}

func (sq *sliceQueue) Pop() *Item {
	sq.lock.Lock()
	popItem := sq.items[0]
	sq.items = sq.items[1:]
	sq.lock.Unlock()
	return &popItem
}

func (sq *sliceQueue) IsEmpty() bool {
	sq.lock.Lock()
	isEmpty := len(sq.items) == 0
	sq.lock.Unlock()
	return isEmpty
}

func (sq *sliceQueue) Size() int {
	sq.lock.Lock()
	size := len(sq.items)
	sq.lock.Unlock()
	return size
}

func (sq *sliceQueue) GetFirstItem() *Item {
	sq.lock.Lock()
	fItem := sq.items[0]
	sq.lock.Unlock()
	return &fItem
}
