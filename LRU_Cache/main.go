/*
Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.

Implement the LRUCache class:

LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
int get(int key) Return the value of the key if the key exists, otherwise return -1.
void put(int key, int value) Update the value of the key if the key exists. Otherwise, add the key-value pair to the cache. If the number of keys exceeds the capacity from this operation, evict the least recently used key.
The functions get and put must each run in O(1) average time complexity.
*/

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

package main

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	c int
	l *list.List
	m map[int]*list.Element
}

type pair struct {
	key, value int
}

func main() {
	obj := Constructor(2)
	obj.Put(1, 1)
	fmt.Println(obj.Get(1))
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		c: capacity,
		l: list.New(),
		m: make(map[int]*list.Element),
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.m[key]; !ok {
		return -1
	} else {
		p := this.l.Remove(node)
		e := this.l.PushFront(p)
		this.m[key] = e
		return p.(pair).value
	}
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.m[key]; ok {
		p := this.l.PushFront(pair{key: key, value: value})
		this.m[key] = p
		this.l.Remove(node)
	} else {
		p := this.l.PushFront(pair{key: key, value: value})
		this.m[key] = p
		if len(this.m) == this.c+1 {
			o := this.l.Remove(this.l.Back())
			p := o.(pair)
			delete(this.m, p.key)
		}
	}
}
