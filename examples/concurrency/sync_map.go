package concurrencyexample

import "sync"

func LoadCount(store *sync.Map, key string) (int, bool) {
	value, ok := store.Load(key)
	if !ok {
		return 0, false
	}
	count, ok := value.(int)
	return count, ok
}
