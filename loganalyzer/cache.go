package loganalyzer

type Cache map[string][]byte

func (c Cache) Exist(key string) bool {
	return c[key] != nil
}

func (c Cache) Get(key string) []byte {
	return c[key]
}

func (c Cache) Set(key string, value []byte) {
	c[key] = value
}

func NewCache() *Cache {
	cache := make(Cache)
	return &cache
}
