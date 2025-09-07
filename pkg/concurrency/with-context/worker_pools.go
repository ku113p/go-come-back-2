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

func doJob(worker uint, job uint) string {
	time.Sleep(100 * time.Millisecond)
	r := fmt.Sprintf("worker #%d | job #%d", worker, job)
	return r
}

func consumer(ctx context.Context, lable uint, jobs <-chan uint, results chan<- string) {
	for {
		select {
		case j, ok := <-jobs:
			if !ok {
				return
			}
			r := doJob(lable, j)
			results <- r
		case <-ctx.Done():
			return
		}
	}
}

func WpExample() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	jobs := make(chan uint)
	results := make(chan string)
	numConsumers := uint(3)
	numJobs := uint(100)

	var wg sync.WaitGroup
	for label := range numConsumers {
		wg.Go(func() {
			consumer(ctx, label, jobs, results)
		})
	}
	go func() {
		wg.Wait()
		close(results)
	}()

	go func() {
		defer close(jobs)
		for j := range numJobs {
			select {
			case jobs <- j:
			case <-ctx.Done():
				return
			}
		}
	}()

	for r := range results {
		fmt.Println(r)
	}
}
