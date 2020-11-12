package gotools

import (
	"sync"
	"time"
)

type memItem struct {
	item                interface{}
	lastUpdatedTime     time.Time
	expiredIntervalSecs int64 //0 means never being expired
}

// MemCache is the simple memory cache
type MemCache struct {
	cached map[string]*memItem
	rwLock *sync.RWMutex
}

func (m *MemCache) startCleanExpiredItem() {
	go func() {
		for {
			m.rwLock.Lock()
			now := time.Now()
			for k, v := range m.cached {
				if now.After(v.lastUpdatedTime.Add(time.Second * time.Duration(v.expiredIntervalSecs))) {
					m.cached[k] = nil
					delete(m.cached, k)
				}
			}
			m.rwLock.Unlock()
			time.Sleep(300 * time.Second)
		}
	}()
}

// NewMemCache is to create a memory cache
func NewMemCache() *MemCache {
	mc := &MemCache{
		cached: make(map[string]*memItem),
		rwLock: new(sync.RWMutex),
	}
	mc.startCleanExpiredItem()
	return mc
}

// Put is to put the item to memory cache
// expiredIntervalSecs is the expired duration (seconds)
func (m *MemCache) Put(key string, item interface{}, expiredIntervalSecs int64) {
	m.rwLock.Lock()
	defer m.rwLock.Unlock()
	mItem := &memItem{
		item:                item,
		expiredIntervalSecs: expiredIntervalSecs,
		lastUpdatedTime:     time.Now(),
	}
	m.cached[key] = mItem
}

// Get the item from the cache
// bool return value means: true = existing, false = not existing/expired
func (m *MemCache) Get(key string) (interface{}, bool) {
	m.rwLock.RLock()
	defer m.rwLock.RUnlock()
	mItem, ok := m.cached[key]
	if !ok {
		return nil, false
	}
	if mItem.expiredIntervalSecs == 0 {
		return mItem.item, true
	}
	if time.Now().After(mItem.lastUpdatedTime.Add(time.Second * time.Duration(mItem.expiredIntervalSecs))) {
		delete(m.cached, key)
		return nil, false
	}
	return mItem.item, true
}

func (m *MemCache) Delete(key string) {
	m.rwLock.Lock()
	defer m.rwLock.Unlock()
	m.cached[key] = nil
	delete(m.cached, key)

}
