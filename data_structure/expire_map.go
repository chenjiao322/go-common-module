package data_structure

import (
	"sync"
	"time"
)

// 这个数据结构是线程安全的
// 锁的用法比较粗暴,有优化空间

type Value struct {
	data     interface{}
	expireAt time.Time
}

type ExpireMap struct {
	data map[string]*Value
	lock sync.Mutex
}

func NewExpireMap() *ExpireMap {
	return &ExpireMap{data: make(map[string]*Value, 0), lock: sync.Mutex{}}
}

func (e *ExpireMap) Add(k string, v interface{}, expire time.Duration) {
	if e == nil {
		return
	}
	e.lock.Lock()
	defer e.lock.Unlock()
	e.data[k] = &Value{
		data:     v,
		expireAt: time.Now().Add(expire),
	}
}

func (e *ExpireMap) Get(k string) (interface{}, bool) {
	if e == nil {
		return nil, false
	}
	e.lock.Lock()
	defer e.lock.Unlock()
	if v, ok := e.data[k]; ok {
		if v.expireAt.Before(time.Now()) {
			delete(e.data, k) // 循环中修改map? 好像没问题.
			return nil, false
		}
		return v.data, true
	}
	return nil, false
}

func (e *ExpireMap) Remove(k string) {
	if e == nil {
		return
	}
	e.lock.Lock()
	defer e.lock.Unlock()
	delete(e.data, k)
}
