/*
 * Copyright 2018 Charles. All Rights Reserved.
 * MAYMOE
 */
package lrucache

import "container/list"

type entry struct {
	key, value interface{}
}

type LRUCache struct {
	max      int
	list     *list.List
	entrys   map[interface{}]*list.Element
	callback func(key, value interface{})
}

func New(limit int) *LRUCache {
	if limit <= 0 {
		return nil
	}
	var lc LRUCache
	lc.max = limit
	lc.list = list.New()
	lc.entrys = make(map[interface{}]*list.Element)
	return &lc
}

// callback after eliminated element
func (lc *LRUCache) OnPop(exec func(key, value interface{})) {
	lc.callback = exec
}

func (lc *LRUCache) Set(key, value interface{}) {
	if lc.list == nil || key == nil || value == nil {
		return
	}
	if el, ok := lc.entrys[key]; ok {
		lc.list.MoveToFront(el)
		el.Value.(*entry).value = value
	} else {
		el = lc.list.PushFront(&entry{key, value})
		lc.entrys[key] = el
		if lc.list.Len() <= lc.max {
			return
		}
		if el = lc.list.Back(); el != nil {
			lc.list.Remove(el)
			e := el.Value.(*entry)
			delete(lc.entrys, e.key)
			if lc.callback != nil {
				lc.callback(e.key, e.value)
			}
		}
	}
}

func (lc *LRUCache) Get(key interface{}) interface{} {
	if lc.list == nil || key == nil {
		return nil
	}
	if el, ok := lc.entrys[key]; ok {
		lc.list.MoveToFront(el)
		return el.Value.(*entry).value
	} else {
		return nil
	}
}

func (lc *LRUCache) Len() int {
	return lc.list.Len()
}

func (lc *LRUCache) Limit() int {
	return lc.max
}
