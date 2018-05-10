package loganalyzer

import "sync"

type Cache struct {
	cache map[string][]byte
	mutex *sync.RWMutex
}

func (c Cache) Exist(key string) bool {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.cache[key] != nil
}

func (c Cache) Get(key string) []byte {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.cache[key]
}

func (c Cache) Set(key string, value []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cache[key] = value
}

func NewCache() *Cache {
	cache := &Cache{
		cache: make(map[string][]byte),
		mutex: new(sync.RWMutex),
	}
	return cache
}
