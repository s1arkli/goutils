package syncx

import "sync"

type Map[K comparable, V any] struct {
	m sync.Map
}

func NewSyncMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		m: sync.Map{},
	}
}

func (m *Map[K, V]) Clear() {
	m.m.Clear()
}

func (m *Map[K, V]) Store(key K, value V) {
	m.m.Store(key, value)
}

func (m *Map[K, V]) Load(key K) (value V, ok bool) {
	v, ok := m.m.Load(key)
	if ok {
		return v.(V), true
	}
	return
}

func (m *Map[K, V]) Range(f func(key K, value V) bool) {
	m.m.Range(func(key, value any) bool {
		return f(key.(K), value.(V))
	})
}

func (m *Map[K, V]) Delete(key K) {
	m.m.Delete(key)
}
