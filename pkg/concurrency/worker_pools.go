package concurrency

import "fmt"

type job struct {
	name string
}
type jobResult struct {
	jobInd string
	worker string
}

func doJob(worker string, j *job) *jobResult {
	r := jobResult{j.name, worker}
	return &r
}

func worker(label string, jobs <-chan *job, results chan<- *jobResult) {
	for j := range jobs {
		r := doJob(label, j)
		results <- r
	}
}

func WpExample() {
	numJobs := 10
	jobs := make(chan *job)
	results := make(chan *jobResult)

	for i := range uint(3) {
		label := fmt.Sprintf("worker #%d", i+1)
		go worker(label, jobs, results)
	}

	go func() {
		defer close(jobs)
		for j := range numJobs {
			name := fmt.Sprintf("job #%d", j+1)
			jobs <- &job{name}
		}
	}()

	for range numJobs {
		r := <-results
		fmt.Println(*r)
	}
}
