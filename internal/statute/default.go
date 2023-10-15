package statute

import (
	"fmt"
	"github.com/bepass-org/dnsutils/internal/cache"
	"net/http"
	"sync"
	"time"
)

// default ttl

const DefaultTTL = 60

// default http client

func DefaultHTTPClient() *http.Client {
	return &http.Client{
		Timeout: 10 * time.Second,
	}
}

// default logger

type Logger interface {
	Debug(s string, v ...interface{})
	Error(s string, v ...interface{})
}

type DefaultLogger struct{}

func (l DefaultLogger) Debug(s string, v ...interface{}) {
	fmt.Printf(fmt.Sprintf("%s\r\n", s), v...)
}

func (l DefaultLogger) Error(s string, v ...interface{}) {
	fmt.Printf(fmt.Sprintf("%s\r\n", s), v...)
}

// default cache

type Cache interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, bool)
}

type DefaultCache struct {
	co   *cache.Cache
	once sync.Once
}

func (c *DefaultCache) prepareCache() {
	c.once.Do(func() {
		c.co = cache.NewCache(DefaultTTL * time.Minute)
	})
}

func (c *DefaultCache) Set(key string, value interface{}) {
	c.prepareCache()
	c.co.Set(key, value)
}

func (c *DefaultCache) Get(key string) (interface{}, bool) {
	c.prepareCache()
	return c.co.Get(key)
}
