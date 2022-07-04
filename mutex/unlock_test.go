package mutex

import (
	"sync"
	"testing"
	"time"
)

func TestUnlock(t *testing.T) {
	var m sync.Mutex
	go func() {
		m.Lock()
	}()
	time.Sleep(1000)
	m.Unlock()
}
