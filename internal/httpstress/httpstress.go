package httpstress

import (
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

type HttpStresser struct {
}

type Job struct {
	ID  int
	Url string
}

type Result struct {
	Job        Job
	StatusCode int
	WorkerID   int
}

var (
	jobs      = make(chan Job)
	resultJob = make(chan Result)
	count     atomic.Int64
	results   = make(map[int]int)
)

func NewHttpStresser() *HttpStresser {
	return &HttpStresser{}
}

func (hS *HttpStresser) StartTest(url string, requestCount, concurrencyCount int) error {

	start := time.Now()
	go allocJob(url, requestCount)
	done := make(chan bool)
	go resultWorker(done)
	createWorkerPool(concurrencyCount)
	<-done
	duration := time.Since(start)

	fmt.Printf("\n\n############ RESULTS ############\n\n")

	for key, value := range results {
		if key == 0 {
			defer fmt.Printf("Total number of error calls: %d\n\n", value)
			continue
		}
		fmt.Printf("Status %d, total responses: %d\n", key, value)
	}
	fmt.Printf("\nTest execution time: %.2f\n\n", duration.Seconds())
	fmt.Printf("Number of requests: %d\n\n", count.Load())

	return nil
}

func createWorkerPool(numberPool int) {
	wg := sync.WaitGroup{}
	wg.Add(numberPool)
	for i := 0; i < int(numberPool); i++ {
		go worker(&wg, i+1)
	}
	wg.Wait()
	close(resultJob)
}

func callHttp(url string) (int, error) {

	req, err := http.Get(url)
	if err != nil && req != nil {
		return req.StatusCode, err
	}

	if err != nil {
		return 0, err
	}

	defer req.Body.Close()

	return req.StatusCode, nil
}

func worker(wg *sync.WaitGroup, ID int) {
	for job := range jobs {
		statusCode, err := callHttp(job.Url)
		if err != nil {
			fmt.Println(err)
		}
		out := Result{job, statusCode, ID}
		resultJob <- out
		count.Add(1)
	}
	wg.Done()
}

func resultWorker(done chan bool) {
	for r := range resultJob {
		fmt.Printf("Worker ID: %d, Job ID: %d, status code: %d, url: %s\n", r.WorkerID, r.Job.ID, r.StatusCode, r.Job.Url)
		_, ok := results[r.StatusCode]
		if !ok {
			results[r.StatusCode] = 1
			continue
		}
		results[r.StatusCode]++
	}
	done <- true
}

func allocJob(url string, requestCount int) {
	for i := 0; i < requestCount; i++ {
		jobs <- Job{i + 1, url}
	}
	close(jobs)
}
