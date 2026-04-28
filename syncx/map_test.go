package syncx

import "testing"

func TestSyncxMap(t *testing.T) {
	m := NewSyncMap[int, int]()
	defer m.Clear()

	//测试正常存值是否可以取出来
	m.Store(1, 1)
	v, ok := m.Load(1)
	if !ok || v != 1 {
		t.Errorf("expect 1 but got %v", v)
	}

	//测试取一个不存在的值是否正常
	v, ok = m.Load(2)
	if ok {
		t.Errorf("expect nil but got %v", v)
	}

	//测试是否能正常删除
	m.Delete(1)
	v, ok = m.Load(1)
	if ok {
		t.Errorf("expect nil but got %v", v)
	}

	count := 0
	m.Store(1, 1)
	m.Store(2, 1)
	m.Store(3, 1)
	m.Range(func(key int, value int) bool {
		if value == 1 {
			count++
		}
		return true
	})
	if count != 3 {
		t.Errorf("expect 3 but got %v", count)
	}
}
