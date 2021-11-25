package gocachebenchmarks

import (
	"sync"
	"time"
)

type Watcher struct {
	fingerprintCacheMap     map[string]*Entry
	fingerprintCacheMapLock sync.RWMutex
	// We might not need expiration -- this is an extra precaution in case fsnotify drops some file notifications
	fingerprintCacheExpirationNano int64
}

type Entry struct {
	val               string
	TimestampUnixNano int64
}

func CreateWatcher() (watcher *Watcher, err error) {
	watcher = &Watcher{
		fingerprintCacheMap:            make(map[string]*Entry),
		fingerprintCacheExpirationNano: (time.Hour * 2).Nanoseconds(),
	}
	return watcher, nil
}

func (w *Watcher) Get(key string) (string, bool) {
	w.fingerprintCacheMapLock.RLock()
	defer w.fingerprintCacheMapLock.RUnlock()

	entry, ok := w.fingerprintCacheMap[key]
	if !ok {
		return "", false
	}
	if time.Now().UnixNano()-entry.TimestampUnixNano < w.fingerprintCacheExpirationNano {
		return entry.val, false
	}
	return "", false
}

func (w *Watcher) Set(key string, val string) {
	w.fingerprintCacheMapLock.Lock()
	defer w.fingerprintCacheMapLock.Unlock()

	w.fingerprintCacheMap[key] = &Entry{val: val, TimestampUnixNano: time.Now().UnixNano()}
}
