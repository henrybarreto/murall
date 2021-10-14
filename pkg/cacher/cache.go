package cacher

// Cache is a simplest way to cache the response
type Cache struct {
	Status bool
	Data  string
}

func (c *Cache) EnableCache(data string) {
	c.Status = true
	c.Data = data
}

func (c *Cache) DisableCache() {
	c.Status = false
	c.Data = ""
}
