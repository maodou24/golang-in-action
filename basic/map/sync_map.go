package _map

import "sync"

type syncMapUsage struct {
	m sync.Map
}

func (s *syncMapUsage) Put(key, value string) {
	s.m.Store(key, value)
}

func (s *syncMapUsage) Get(key string) (string, bool) {
	value, ok := s.m.Load(key)
	if !ok {
		return "", false
	}

	str, ok := value.(string)
	return str, ok
}

func (s *syncMapUsage) Delete(key string) {
	s.m.Delete(key)
}

func (s *syncMapUsage) DeleteValue(key, value string) bool {
	return s.m.CompareAndDelete(key, value)
}

func (s *syncMapUsage) Keys() []string {
	var keys []string
	s.m.Range(func(key, _ any) bool {
		if str, ok := key.(string); ok {
			keys = append(keys, str)
		}
		return true
	})
	return keys
}

func (s *syncMapUsage) Values() []string {
	var keys []string
	s.m.Range(func(_, value any) bool {
		if str, ok := value.(string); ok {
			keys = append(keys, str)
		}
		return true
	})
	return keys
}

func (s *syncMapUsage) Clear() {
	s.m.Clear()
}

