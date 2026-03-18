package crawler

import "sync"

type SafeSet struct {
	visited map[string]struct{}
	mutex   sync.RWMutex
}

func NewSafeSet() *SafeSet {
	return &SafeSet{visited: make(map[string]struct{})}
}

func (s *SafeSet) Add(url string) bool {

	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, exists := s.visited[url]; exists {
		return false
	}

	s.visited[url] = struct{}{}

	return true
}
