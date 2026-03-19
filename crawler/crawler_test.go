package crawler

import (
	"sync"
	"testing"
)

func TestSafeSet_Add(t *testing.T) {
	cache := NewSafeSet()

	if !cache.Add("http://stress-test.com") {
		t.Error("(TEST) >> expected true for new URL")
	}

	if cache.Add("http://stress-test.com") {
		t.Error("(TEST) >> expected false for duplicate URL")
	}
}

func TestSafeSet_Concurrency(t *testing.T) {
	cache := NewSafeSet()
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cache.Add("http://stress-test.com")
		}()
	}

	wg.Wait()
}
