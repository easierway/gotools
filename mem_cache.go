package aws_plugins

import (
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
}

// NewMemCache is to create a memory cache
func NewMemCache() *MemCache {
	return &MemCache{
		cached: make(map[string]*memItem),
	}
}

// Put is to put the item to memory cache
// expiredIntervalSecs is the expired duration (seconds)
func (m *MemCache) Put(key string, item interface{}, expiredIntervalSecs int64) {
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
	mItem, ok := m.cached[key]
	if !ok {
		return nil, false
	}
	if time.Now().After(mItem.lastUpdatedTime.Add(time.Second * time.Duration(mItem.expiredIntervalSecs))) {
		return nil, false
	}
	return mItem.item, true
}
