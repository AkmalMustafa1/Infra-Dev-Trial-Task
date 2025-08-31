package ratelimit

import (
    "sync"
    "time"
)

type client struct {
    requests int
    last     time.Time
}

var (
    clients = map[string]*client{}
    mu      sync.Mutex
    limit   = 10
    window  = time.Minute
)

func Allow(ip string) bool {
    mu.Lock()
    defer mu.Unlock()

    c, ok := clients[ip]
    now := time.Now()
    if !ok || now.Sub(c.last) > window {
        clients[ip] = &client{requests: 1, last: now}
        return true
    }

    if c.requests >= limit {
        return false
    }

    c.requests++
    c.last = now
    return true
}
