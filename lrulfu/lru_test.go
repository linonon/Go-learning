package main

import "testing"

func TestLruCache(t *testing.T) {
	cache := NewLruCache[int, string](3)

	cache.Put(1, "L")
	cache.PrintAll()

	cache.Put(2, "R")
	cache.Put(3, "U")
	cache.PrintAll()

	cache.Get(2)
	cache.PrintAll()

	cache.Get(1)
	cache.PrintAll()

	cache.Put(4, "GwG")
	cache.PrintAll()
}
