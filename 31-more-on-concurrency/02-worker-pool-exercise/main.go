package mai

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id int
	randomno int
}

type Result struct {
	job Job
	sumofdigits, wid int
}

func digits(number int) int {
	sum := 0
	no := number
	for no !=0 {
		num := no % 10
		sum = sum + num
		no = no / 10
	}
	time.Sleep(2 * time.Second)
	return sum
}

func allocateJobs(noofjobs int, jobs chan Job) {
	for i:=0; i < noofjobs ; i++ {
		randomno := rand.Intn(999) + 1
		job := Job{i, randomno}
		jobs <- job
	} 
	close(jobs)
}

func createWorkerPool(noOfWorkers int, jobs chan Job, results chan Result) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for job := range jobs {
				output := Result{job, digits(job.randomno), i}
				results <- output
			}
		}(i)
	}
	wg.Wait()
	close(results)
}

func main() {
	jobs := make(chan Job, 10)
	results := make(chan Result, 10)
	noOfJobs := 50
	noOfWorkers := 10
	go allocateJobs(noOfJobs, jobs)
	go createWorkerPool(noOfWorkers, jobs, results)
	for result := range results {
        fmt.Printf("Job id %d, input random no %d, sum of digits %d, worker id %d\n", result.job.id, result.job.randomno, result.sumofdigits, result.wId)
    }
}