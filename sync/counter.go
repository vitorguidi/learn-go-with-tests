package sync

type Counter struct {
	value int
}

func (c *Counter) Increment() {
	c.value += 1
}

func (c *Counter) Value() int {
	return c.value
}

func NewCounter() *Counter {
	return &Counter{value: 0}
}
