package main

import (
	"fmt"
)

// LFU (Least Frequently Used)
//
//	核心是使用 LinkedHashMap 鏈式哈希表
//	除了正常情況下會淘汰次數最少的節點外, 假如擁有多個同樣使用次數的節點, 則會淘汰最舊的那個
//
// 主要規定:
//	1. get(key), 返回 val
//  2. get 或 put 過的 key, 這個 key 的 freq 都需要 +1
//  3. 滿了的時候再加入, 需要將 freq 最小的 key 刪除, 如果最小的 freq 有多個 key, 刪除最舊的那個

type LfuCache[K, V comparable] struct {
	// ketToVal 存儲 key 到 val 的映射, 用於 get(key)
	keyToVal map[K]V
	// keyToFreq 存儲 key 的使用次數
	keyToFreq map[K]int

	// minFreq 找到 freq 最小的 key, 用 minFreq 記錄, 用於 O(1) 級的刪除元素
	minFreq int
	// freqToKeys 多個 key 會用相同的 freq, 所以 "freq 對 key" 是一對多的關係,
	// "freq to key" 用雙向鏈錶, 這樣方便刪除最舊的 key
	freqToKeys map[int]*LruCache[K, V]

	capacity int
}

func NewLfuCache[K, V comparable](cap int) *LfuCache[K, V] {
	return &LfuCache[K, V]{
		keyToVal:  make(map[K]V),
		keyToFreq: make(map[K]int),

		minFreq:    0,
		freqToKeys: make(map[int]*LruCache[K, V]),

		capacity: cap,
	}
}

// Get 取一個數, 並更新 freq
func (c *LfuCache[K, V]) Get(key K) (V, bool) {
	v, ok := c.keyToVal[key]
	if !ok {
		return v, false
	}

	// increase freq
	c.increaseFreq(key, v)
	return v, true
}

func (c *LfuCache[K, V]) Put(key K, value V) {
	if c.capacity <= 0 {
		return
	}

	// 如果 key 已存在, 更新次數 keyToFreq, 更新位置 freqToKeys
	_, ok := c.keyToVal[key]
	if ok {
		c.keyToVal[key] = value
		c.increaseFreq(key, value)
		return
	}

	// if key not exist, put
	if len(c.keyToVal) == c.capacity {
		// remove minFreq or oldest key
		c.removeMinFreq()
	}

	c.keyToVal[key] = value
	c.keyToFreq[key] = 1
	c.minFreq = 1
	c.putFreqToKey(1, key, value)
}

func (c *LfuCache[K, V]) putFreqToKey(freq int, key K, value V) {
	if _, ok := c.freqToKeys[freq]; !ok {
		c.freqToKeys[freq] = NewLruCache[K, V](-1)
	}
	c.freqToKeys[freq].Put(key, value)
}

func (c *LfuCache[K, V]) increaseFreq(key K, value V) {
	freq := c.keyToFreq[key]
	fmt.Println(freq)
	// key 的使用次數 +1
	c.keyToFreq[key]++
	// freqToKey 從舊的刪除, 加到新的
	c.freqToKeys[freq].RemoveNode(key)
	c.putFreqToKey(freq+1, key, value)
	// 如果舊的 freqToKey 為空, 刪除這個 key
	if c.freqToKeys[freq].Size() == 0 {
		delete(c.freqToKeys, freq)
		if freq == c.minFreq {
			c.minFreq++
		}
	}
}

func (c *LfuCache[K, V]) removeMinFreq() {
	minLru := c.freqToKeys[c.minFreq]
	lastNode := minLru.RemoveLastNode()

	if minLru.Size() == 0 {
		delete(c.freqToKeys, c.minFreq)
		// 不需要更新 c.minFreq, 因為新增數據總是從 1 開始
	}

	delete(c.keyToVal, lastNode.Key)
	delete(c.keyToFreq, lastNode.Key)
}

func (c *LfuCache[K, V]) PrintAll() {
	fmt.Printf("%+v\n", c)
}
