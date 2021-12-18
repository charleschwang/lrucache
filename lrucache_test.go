/*
 * Copyright (c) 2021, CharlesChwang.
 * All rights reserved.
 *
 * @Date: 2021-12-19 02:54:06
 * @LastEditTime: 2021-12-19 04:03:42
 * @FilePath: \lrucache\lrucache_test.go
 */
package lrucache

import "testing"

func TestSet(t *testing.T) {
	cache := New(3)
	cache.Set(1, "a")
	cache.Set(2, "b")
	if cache.Len() != 2 {
		t.Fatal("Set test fail")
	}
	if cache.Get(2) != "b" {
		t.Fatal("Set test fail")
	}
}

func TestGet(t *testing.T) {
	cache := New(3)
	cache.Set("a", 1)
	cache.Set("b", "2")
	if v := cache.Get("a"); v != 1 {
		t.Fatal("Get test fail")
	}
	if v := cache.Get("b"); v != "2" {
		t.Fatal("Get test fail")
	}
	if cache.Get("c") != nil {
		t.Fatal("Get test fail")
	}
}

func TestLRU(t *testing.T) {
	cache := New(5)
	cache.Set("a", 1)
	cache.Set("b", 2)
	cache.Set("c", 3)
	cache.Set("d", 4)
	cache.Set("e", 5)
	cache.Get("e")
	cache.Get("d")
	cache.Get("e")
	// should evict "a", "b" and "c"
	num := 0
	cache.OnPop(func(key, value interface{}) {
		num += value.(int)
	})
	cache.Set("f", 6)
	cache.Set("i", 7)
	cache.Set("j", 8)
	if num != 1+2+3 {
		t.Fatal("The LRU strategy doesn't work")
	}
	if cache.Get("a") != nil {
		t.Fatal("The LRU strategy doesn't work")
	}
	if cache.Get("e") == nil {
		t.Fatal("The LRU strategy doesn't work")
	}
}
