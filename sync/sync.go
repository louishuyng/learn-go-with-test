package sync

import "sync"

func NewCounter() *Counter {
	return &Counter{}
}

type Counter struct {
	value int
	mu    sync.Mutex
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
