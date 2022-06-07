package cache

import "time"

type Data struct {
	val      string
	deadline time.Time
}

type Cache struct {
	m map[string]Data
}

func NewCache() *Cache {
	return &Cache{
		make(map[string]Data),
	}
}

func (c *Cache) Get(key string) (string, bool) {
	startTime := time.Now()

	data, ok := c.m[key]
	if !ok {
		return "", ok
	}

	if !data.deadline.IsZero() {
		notExp := startTime.Before(data.deadline)
		if notExp {
			return data.val, true
		} else {
			delete(c.m, key)
			return "", false
		}
	}
	return data.val, ok
}

func (c *Cache) Keys() []string {
	var arr []string
	for key := range c.m {
		arr = append(arr, key)
	}
	return arr
}

func (c *Cache) Put(key, value string) {
	data := Data{
		val:      value,
		deadline: time.Time{},
	}
	c.m[key] = data
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	data := Data{
		val:      value,
		deadline: deadline,
	}
	c.m[key] = data
}
