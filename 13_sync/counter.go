package counter

import "sync"

type Counter struct {
	mut   sync.Mutex
	value int
}

func (c *Counter) Inc() {
    c.mut.Lock()
    defer c.mut.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}

func NewCounter() *Counter {
    return &Counter{}
}
