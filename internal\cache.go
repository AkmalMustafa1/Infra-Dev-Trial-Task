package cache

import (
    "sync"
    "time"
)

type cacheItem struct {
    value float64
    ttl   time.Time
}

var (
    store = make(map[string]cacheItem)
    mu    sync.Mutex
)

func GetOrFetch(key string, fetch func() (float64, error)) (float64, error) {
    mu.Lock()
    item, exists := store[key]
    mu.Unlock()

    if exists && time.Now().Before(item.ttl) {
        return item.value, nil
    }

    // Fetch new value and store
    mu.Lock()
    defer mu.Unlock()
    val, err := fetch()
    if err != nil {
        return 0, err
    }
    store[key] = cacheItem{value: val, ttl: time.Now().Add(10 * time.Second)}
    return val, nil
}
