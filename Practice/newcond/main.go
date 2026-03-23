package main

import (
	"fmt"
	"sync"
	"time"
)

const bufferSize = 5

type Buffer struct {
	items []int
	mu    sync.Mutex
	cond  *sync.Cond
}

func newBuffer(size int) *Buffer {
	b := &Buffer{
		items: make([]int, 0, size),
	}
	b.cond = sync.NewCond(&b.mu)
	return b
}

func (b *Buffer) produce(item int) {
	b.mu.Lock()
	defer b.mu.Unlock()
	for len(b.items) == bufferSize {
		b.cond.Wait()
	}
	b.items = append(b.items, item)
	fmt.Println("Produced:", item)
	b.cond.Signal()
}

func (b *Buffer) consume() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	for len(b.items) == 0 {
		b.cond.Wait()
		//Waiting to be appended
	}
	item := b.items[0]
	b.items = b.items[1:]
	fmt.Println("Consumed:", item)
	b.cond.Signal()
	return item
}

func producer(b *Buffer, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range 10 {
		b.produce(i + 100)
		time.Sleep(100 * time.Millisecond)
	}
}

func consumer(b *Buffer, wg *sync.WaitGroup) {
	defer wg.Done()
	for range 10 {
		b.consume()
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	buffer := newBuffer(bufferSize)
	var wg sync.WaitGroup
	wg.Add(2)
	go producer(buffer, &wg)
	go consumer(buffer, &wg)

	wg.Wait()
}
