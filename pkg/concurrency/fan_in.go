package concurrency

import (
	"fmt"
	"sync"
	"time"
)

type item struct {
	worker  uint
	counter uint
}

func producer(index uint) chan *item {
	ch := make(chan *item)

	go func() {
		defer close(ch)
		for i := range uint(10) {
			obj := item{index, i}
			ch <- &obj
			time.Sleep(50 * time.Millisecond)
		}
	}()

	return ch
}

func fanIn(inputs []chan *item) <-chan *item {
	out := make(chan *item)

	go func() {
		defer close(out)

		var wg sync.WaitGroup
		defer wg.Wait()

		for _, ch := range inputs {
			wg.Go(func() {
				for i := range ch {
					out <- i
				}
			})
		}
	}()

	return out
}

func FiExample() {
	var inp_chans []chan *item

	for i := range uint(5) {
		ch := producer(i)
		inp_chans = append(inp_chans, ch)
	}

	collector := fanIn(inp_chans)

	for c := range collector {
		fmt.Printf("%v\n", *c)
	}
}
