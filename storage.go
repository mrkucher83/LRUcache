package LRUcache

import "sync"

type LRUCache struct {
	mu       sync.RWMutex
	capacity int
	index    map[string]*Node
	storage  *List
}

func NewCache(n int) *LRUCache {
	return &LRUCache{
		capacity: n,
		index:    make(map[string]*Node, n),
		storage:  NewList(),
	}
}

func (c *LRUCache) Add(key, value string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.index[key]; ok {
		return false
	}

	c.index[key] = c.storage.PushFront(value)

	if c.storage.Len() > c.capacity {
		delete(c.index, key)
		last := c.storage.Last()
		c.storage.Remove(last)
	}
	return true
}

func (c *LRUCache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if node, ok := c.index[key]; ok {
		c.storage.MoveToFront(node)
		return node.Val.(string), true
	}

	return "", false
}

func (c *LRUCache) Remove(key string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if node, ok := c.index[key]; ok {
		delete(c.index, key)
		c.storage.Remove(node)
		return true
	}

	return false
}

func (c *LRUCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.index = nil
	c.storage = NewList()
}
