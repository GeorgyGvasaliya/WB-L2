package pkg

type evictionAlgo interface {
	evict(c *Cache)
}
