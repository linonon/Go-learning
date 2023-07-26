package main

import "fmt"

// LRU 核心
//
//	Get Data: 如果命中了 Cache 會返回資料前, 將資料放到 Head, 防止下次再次命中時, 資料被淘汰
// 	Put Data: 先確認是否存在資料
//		如果存在, 更新資料, 並將資料放到 Head, 防止下次再次命中時, 資料被淘汰
//		不存在, 將資料放到 Head
//
//	需要實現兩個數據結構: 雙向鏈表, HashMap
//	雙向鏈表: 用於 O(1) 級的移動數據, 並且可以快速的將資料放到 Head
//	HashMap: 用於 O(1) 查找數據是否存在, 並且可以快速的將資料放到 Head

// Node 節點
type Node[K, V comparable] struct {
	Prev  *Node[K, V]
	Next  *Node[K, V]
	Key   K
	Value V
}

func NewNode[K, V comparable](key K, value V) *Node[K, V] {
	return &Node[K, V]{
		Prev:  nil,
		Next:  nil,
		Key:   key,
		Value: value,
	}
}

// /////////////////////////////////////
// DoubleLinkedList 雙向鏈錶
// /////////////////////////////////////
type DoubleLinkedList[K, V comparable] struct {
	Head *Node[K, V]
	Tail *Node[K, V]
}

func NewDoubleLinkedList[K, V comparable]() *DoubleLinkedList[K, V] {
	head := &Node[K, V]{}
	tail := &Node[K, V]{}

	head.Next = tail
	tail.Prev = head
	return &DoubleLinkedList[K, V]{
		Head: head,
		Tail: tail,
	}
}

func (dll *DoubleLinkedList[K, V]) AddHead(node *Node[K, V]) {
	node.Next = dll.Head.Next
	node.Prev = dll.Head

	dll.Head.Next.Prev = node
	dll.Head.Next = node
}

func (dll *DoubleLinkedList[K, V]) RemoveNode(node *Node[K, V]) {
	node.Next.Prev = node.Prev
	node.Prev.Next = node.Next

	node.Next, node.Prev = nil, nil
}

func (dll *DoubleLinkedList[K, V]) GetLastNode() *Node[K, V] {
	return dll.Tail.Prev
}

func (dll *DoubleLinkedList[K, V]) PrintAll() {
	fmt.Println("\n Print all node")
	cur := dll.Head.Next
	for cur != dll.Tail {
		fmt.Printf("key: %+v, value:%+v\n ", cur.Key, cur.Value)
		cur = cur.Next
	}
}

// /////////////////////////////////////
// LruCache LRU Algorithm
// /////////////////////////////////////
type LruCache[K, V comparable] struct {
	nodeMap  map[K]*Node[K, V]
	list     *DoubleLinkedList[K, V]
	capacity int // capacity < 0 的話, 表示不會自動淘汰
}

func NewLruCache[K, V comparable](cap int) *LruCache[K, V] {
	return &LruCache[K, V]{
		nodeMap:  make(map[K]*Node[K, V]),
		list:     NewDoubleLinkedList[K, V](),
		capacity: cap,
	}
}

// Get 取得對應的 node, 然後更新 node 的位置
func (c *LruCache[K, V]) Get(key K) (Node[K, V], bool) {
	v, ok := c.nodeMap[key]
	if !ok {
		return *v, false
	}

	c.list.RemoveNode(v)
	c.list.AddHead(v)
	return *v, true
}

// Put 放 node, 如果超過 cap, 刪掉 List 裡最後一個, 並且將 map 裡的 key 也刪除後,
// 再將 newNode 放進 List 和 Map 中
func (c *LruCache[K, V]) Put(key K, value V) {
	newNode := &Node[K, V]{
		Prev:  nil,
		Next:  nil,
		Key:   key,
		Value: value,
	}

	if node, ok := c.nodeMap[key]; ok {
		c.list.RemoveNode(node)
		c.list.AddHead(newNode)
		c.nodeMap[key] = newNode
	} else {
		if c.capacity >= 0 && len(c.nodeMap) == c.capacity {
			c.RemoveLastNode()
		}

		c.nodeMap[key] = newNode
		c.list.AddHead(newNode)
	}
}

func (c *LruCache[K, V]) Size() int {
	return len(c.nodeMap)
}

// RemoveNode is O(1), will remove specific node
func (c *LruCache[K, V]) RemoveNode(key K) {
	node, ok := c.FindNode(key)
	if !ok {
		return
	}

	c.list.RemoveNode(node)
	delete(c.nodeMap, key)
}

// FindNode is O(1)
func (c *LruCache[K, V]) FindNode(key K) (*Node[K, V], bool) {
	v, ok := c.nodeMap[key]
	return v, ok
}

func (c *LruCache[K, V]) RemoveLastNode() Node[K, V] {
	lastNode := c.list.GetLastNode()
	c.list.RemoveNode(lastNode)
	delete(c.nodeMap, lastNode.Key)

	return *lastNode
}

func (c *LruCache[K, V]) PrintAll() {
	c.list.PrintAll()
}
