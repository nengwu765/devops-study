package main

import "sync"

type Set interface {
	Put(key string, item Item)
	Keys() []string
	Contains(key string) bool
	Remove(key string)

	// PutIfAbsent 如果之前已经有了，就返回旧的值，absent = false
	// 如果之前没有，就塞下去， absent = true
	PutIfAbsent(key string, item Item) (old Item, absent bool)
}

type Item interface {}

var _ Set = &mapSet{}

// New一个MapSet
func NewMapSet() Set {
	return &mapSet{
		items: map[string]Item{},
	}
}

type mapSet struct {
	items map[string]Item
	lock  sync.RWMutex
}

// Put 原对应的key会被覆盖
func (s *mapSet) Put(key string, item Item) {
	s.lock.Lock()
	s.items[key] = item
	s.lock.Unlock()
}

func (s *mapSet) Keys() []string {
	s.lock.Lock()
	var keys []string
	for key, _ := range s.items {
		keys = append(keys, key)
	}
	s.lock.Unlock()
	return keys
}

func (s *mapSet) Contains(key string) bool {
	s.lock.Lock()
	_, ok := s.items[key]
	s.lock.Unlock()
	return ok
}

func (s *mapSet) Remove(key string) {
	delete(s.items, key)
}

func (s *mapSet) PutIfAbsent(key string, item Item) (old Item, absent bool) {
	s.lock.Lock()
	v, ok := s.items[key]
	if !ok {
		s.items[key] = item
	}
	s.lock.Unlock()

	return v, ok
}
