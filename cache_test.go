package gocachebenchmarks

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/coocood/freecache"
	cache "github.com/patrickmn/go-cache"
)

func randomPath() string {
	tenantID := rand.Intn(1000)
	return fmt.Sprintf("/ns/epdlp/tenantpolicy/%d/policy.zip", tenantID)
}

func randomSha1() string {
	s := sha1.New()
	_, _ = s.Write([]byte(strconv.Itoa(rand.Int())))
	return hex.EncodeToString(s.Sum(nil))
}

func BenchmarkGoCache(b *testing.B) {
	c := cache.New(1*time.Minute, 5*time.Minute)

	b.Run("set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c.Add(randomPath(), randomSha1(), cache.DefaultExpiration)
		}
	})

	b.Run("get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			value, found := c.Get(randomPath())
			if found {
				_ = value.(string)
			}
		}
	})
}

func BenchmarkGoCacheParallel(b *testing.B) {
	c := cache.New(1*time.Minute, 5*time.Minute)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.Add(randomPath(), randomSha1(), time.Hour*2)
			value, found := c.Get(randomPath())
			if found {
				_ = value
			}
		}
	})
}

func BenchmarkFreecache(b *testing.B) {
	c := freecache.NewCache(1024 * 1024 * 5)

	b.Run("Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c.Set([]byte(randomPath()), []byte(randomPath()), 60)
		}
	})

	b.Run("Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			value, err := c.Get([]byte(randomPath()))
			if err == nil {
				_ = value
			}
		}
	})
}

func BenchmarkFreecacheParallel(b *testing.B) {
	c := freecache.NewCache(1024 * 1024 * 5)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.Set([]byte(randomPath()), []byte(randomSha1()), 2*60*60)
			value, err := c.Get([]byte(randomPath()))
			if err == nil {
				_ = value
			}
		}
	})
}

func BenchmarkFsWatcher(b *testing.B) {
	c, _ := CreateWatcher()

	b.Run("Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c.Set(randomPath(), randomPath())
		}
	})

	b.Run("Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			value, ok := c.Get(randomPath())
			if ok {
				_ = value
			}
		}
	})
}

func BenchmarkFsWatcherParallel(b *testing.B) {
	c, _ := CreateWatcher()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.Set(randomPath(), randomSha1())
			value, ok := c.Get(randomPath())
			if ok {
				_ = value
			}
		}
	})
}
