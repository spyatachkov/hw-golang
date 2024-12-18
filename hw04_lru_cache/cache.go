package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type Data struct {
	Key   Key
	Value interface{}
}

type lruCache struct {
	mu       sync.Mutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, exist := c.items[key]

	if c.queue.Len() == c.capacity {
		lastElem := c.queue.Back()

		d := c.castData(lastElem.Value)

		c.queue.Remove(lastElem)
		delete(c.items, d.Key)
	}

	if exist {
		if val.Value != value {
			c.queue.Remove(val)
			e := c.queue.PushFront(Data{Key: key, Value: value})
			c.items[key] = e
			return true
		}

		c.queue.MoveToFront(val)
		return true
	}

	e := c.queue.PushFront(Data{Key: key, Value: value})
	c.items[key] = e

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	val, exist := c.items[key]

	if exist {
		c.queue.MoveToFront(val)
		d := c.castData(val.Value)
		return d.Value, true
	}

	return nil, false
}

func (c *lruCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for i := c.queue.Front(); i != nil; i = i.Next {
		i.Prev = nil
		i.Next = nil
	}

	for key := range c.items {
		delete(c.items, key)
	}
}

func (c *lruCache) castData(i interface{}) Data {
	data, ok := i.(Data)
	if ok {
		return Data{Key: data.Key, Value: data.Value}
	}
	panic("cast error")
}
