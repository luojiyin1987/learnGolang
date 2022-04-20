package LRU

import (
	"container/list"
	"errors"
	"sync"
)

type CallBack func(key interface{}, value interface{})

type Lru struct {
	max   int
	l     *list.List
	cb    CallBack
	cache map[interface{}]*list.Element
	mu    *sync.Mutex
}

type Node struct {
	Key interface{}
	Val interface{}
}

func NewLru(len int, c CallBack) *Lru {
	return &Lru{
		max:   len,
		l:     list.New(),
		cb:    c,
		cache: make(map[interface{}]*list.Element),
		mu:    new(sync.Mutex),
	}
}

func (l *Lru) Get(key interface{}) (value interface{}, ok bool) {
	if l.cache == nil {
		return nil, false
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	if e, ok := l.cache[key]; ok {
		l.l.MoveToFront(e)
		return e.Value.(*Node).Val, true
	}

	return nil, false
}

func (l *Lru) GetAll() []*Node {
	l.mu.Lock()
	defer l.mu.Lock()
	var data []*Node
	for _, v := range l.cache {
		data = append(data, v.Value.(*Node))
	}
	return data
}

func (l *Lru) Del(key interface{}) {
	if l.cache == nil {
		return
	}
	l.mu.Lock()
	defer l.mu.Unlock()

	if e, ok := l.cache[key]; ok {
		l.l.Remove(e)
		delete(l.cache, key)

		if l.cb != nil {
			node := e.Value.(*Node)
			l.cb(node.Key, node.Val)
		}
	}
}

func (l *Lru) Add(key interface{}, val interface{}) error {
	if l.l == nil {
		return errors.New("list is nil")
	}
	l.mu.Lock()
	defer l.mu.Unlock()

	if e, ok := l.cache[key]; ok {
		e.Value.(*Node).Val = val
		l.l.MoveToFront(e)
		return nil
	}

	ele := l.l.PushFront(&Node{Key: key, Val: val})
	l.cache[key] = ele

	if l.max != 0 && l.l.Len() > l.max {
		if e := l.l.Back(); e != nil {
			l.l.Remove(e)
			node := e.Value.(*Node)
			delete(l.cache, node.Key)
			if l.cb != nil {
				l.cb(node.Key, node.Val)
			}
		}
	}
	return nil
}
