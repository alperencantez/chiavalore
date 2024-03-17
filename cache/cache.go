package cache

import "sync"

type CV_Cache struct {
	cache map[string]interface{}
	mutex sync.Mutex
}

func InitCV_Cache() *CV_Cache {
	return &CV_Cache{cache: make(map[string]interface{})}
}

func (c *CV_Cache) Set(key string, value interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.cache[key] = value
}

func (c *CV_Cache) Get(key string) interface{} {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.cache[key]
}

func (c *CV_Cache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	delete(c.cache, key)
}
