package concurrent

import "sync"

type ConcurrentSlice struct {
	sync.RWMutex
	items []interface{}
}

func (slice *ConcurrentSlice) Append(item interface{}) {
	slice.Lock()
	defer slice.Unlock()

	slice.items = append(slice.items, item)
}

func (slice *ConcurrentSlice) PopAll() <-chan interface{} {
	c := make(chan interface{})

	f := func() {
		slice.Lock()
		defer slice.Unlock()
		for _, value := range slice.items {
			c <- value
		}

		slice.items = make([]interface{}, 0)
		close(c)
	}
	go f()

	return c
}
