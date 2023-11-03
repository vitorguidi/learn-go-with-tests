package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("Should get 10000 after incrementing in 10000 goroutines", func(t *testing.T) {
		wg := sync.WaitGroup{}
		wg.Add(10000)
		mut := sync.Mutex{}
		c := NewCounter()
		for i := 0; i < 10000; i++ {
			go func() {
				mut.Lock()
				defer wg.Done()
				defer mut.Unlock()
				c.Increment()
			}()
		}
		wg.Wait()
		got := c.value
		if got != 10000 {
			t.Errorf("Wrong counter value: expected 10000, got %d", got)
		}
	})
}
