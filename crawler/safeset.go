package crawler

import "sync"

type SafeSet struct {
	visited map[string]struct{}
	mutex   sync.RWMutex
}

func NewSafeSet(s *SafeSet) *SafeSet {
	return &SafeSet{visited: s.visited}
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
