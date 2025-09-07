package withcontext

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type job struct {
	name string
}

func newJob(producerLabel string, ind uint) *job {
	name := fmt.Sprintf("[%s] task #%d", producerLabel, ind)
	return &job{name}
}

func producer(ctx context.Context, label string, numJobs uint, timeout time.Duration) <-chan *job {
	ch := make(chan *job)

	go func() {
		defer close(ch)

		fst := make(chan any, 1)
		defer close(fst)
		fst <- nil

		ticker := time.NewTicker(timeout)
		for i := range numJobs {
			select {
			case <-fst:
				ch <- newJob(label, i)
			case <-ticker.C:
				ch <- newJob(label, i)
			case <-ctx.Done():
				return
			}
		}
	}()

	return ch
}

func fanIn(ctx context.Context, producers []<-chan *job) <-chan *job {
	ch := make(chan *job)

	go func() {
		defer close(ch)

		var wg sync.WaitGroup
		defer wg.Wait()

		for _, p := range producers {
			wg.Add(1)
			go func(p <-chan *job) {
				defer wg.Done()
				for {
					select {
					case j, ok := <-p:
						if !ok {
							return
						}
						ch <- j
					case <-ctx.Done():
						return
					}
				}
			}(p)
		}
	}()

	return ch
}

func FiExample() {
	var producers []<-chan *job
	numProducers, numJobs := 100, uint(100)
	producerTimeout := 50 * time.Millisecond
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	for ind := range numProducers {
		label := fmt.Sprintf("producer #%d", ind)
		p := producer(ctx, label, numJobs, producerTimeout)
		producers = append(producers, p)
	}

	collector := fanIn(ctx, producers)

	for c := range collector {
		fmt.Printf("%v\n", *c)
	}
}
