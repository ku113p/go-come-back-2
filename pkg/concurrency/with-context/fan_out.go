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

func makeJob(worker uint, job string) {
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("[%d] job %s done\n", worker, job)
}

func worker(ctx context.Context, label uint, jobs chan string) {
	for {
		select {
		case j, ok := <-jobs:
			if !ok {
				return
			}
			makeJob(label, j)
		case <-ctx.Done():
			return
		}
	}
}

func FoExample() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	jobs := make(chan string)
	numWorkers := uint(10)
	numJobs := 100

	var wg sync.WaitGroup
	defer wg.Wait()
	for ind := range numWorkers {
		wg.Go(func() {
			worker(ctx, ind, jobs)
		})
	}

	go func() {
		defer close(jobs)
		for ind := range numJobs {
			job := fmt.Sprintf("#%d", ind)
			select {
			case jobs <- job:
			case <-ctx.Done():
				return
			}
		}
	}()
}
