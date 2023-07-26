package main

import "testing"

func TestLfuCache(t *testing.T) {
	cache := NewLfuCache[int, string](2)

	cache.Put(1, "L")
	cache.Put(2, "F")
	cache.Put(3, "U")

	cache.PrintAll()

	cache.Get(3)
	cache.Get(3)
	cache.Get(3)
	cache.Get(2)
	cache.PrintAll()

	cache.Put(3, "X")
	cache.PrintAll()
	cache.Put(4, "X")
	cache.PrintAll()
}
