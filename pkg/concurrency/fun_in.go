package concurrency

import (
	"fmt"
	"time"
)

func worker(ch <-chan int, index uint) {
	counter := 0
	for w := range ch {
		counter++
		text := fmt.Sprintf("[%d] work name='%v'. total=%d", index, w, counter)
		fmt.Println(text)
		time.Sleep(1 * time.Second)
	}
}

func Example() {
	work := make(chan int)

	for i := range uint(3) {
		go worker(work, i)
	}

	for i := range 10 {
		work <- i
	}
}
