package workerpool

import (
	"fmt"
	"math/rand"
	"sync"
)

func sumdigits(num int) int {
	res := 0
	for num > 0 {
		res += num % 10
		num = num / 10
	}
	return res
}

type Job struct {
	id       int
	randomno int
}

type Result struct {
	job              Job
	sumofdigits, wId int
}

func ProcessSumDigit() {
	var wg sync.WaitGroup
	jobchn := getJobChannel(&wg)
	reschn := calculateSumDigit(jobchn, 10, &wg)
	go func() {
		defer close(reschn)
		for actres := range reschn {
			fmt.Printf("result is=%v\n", actres)
		}
	}()
	wg.Wait()
	fmt.Println("----END------")

}

func getJobChannel(awg *sync.WaitGroup) chan Job {
	jobchan := make(chan Job, 10)
	go func() {
		defer close(jobchan)
		for i := 1; i <= 50; i++ {
			jobchan <- Job{
				id:       i,
				randomno: rand.Intn(1000),
			}
		}
	}()
	return jobchan
}

func calculateSumDigit(jobchn chan Job, numworker int, awg *sync.WaitGroup) chan Result {
	reschn := make(chan Result, 10)
	for i := 0; i < numworker; i++ {
		awg.Add(1)
		go func(wid int) {
			defer awg.Done()
			for thejob := range jobchn {
				sum := sumdigits(thejob.randomno)
				theres := Result{
					job:         thejob,
					sumofdigits: sum,
					wId:         wid,
				}
				reschn <- theres
			}
		}(i)
	}
	return reschn
}
