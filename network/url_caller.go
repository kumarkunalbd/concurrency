package network

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// This is getting blocked after executoion
func MultiURLTime(urls []string) {
	producedChn, errChn := producerChan(urls)
	var wg sync.WaitGroup
	fetcherChn := dataFetcher(producedChn, errChn, &wg)
	counter := 0
NEXTEXECUTION:
	for {
		select {
		case respRes, open := <-fetcherChn:
			if !open {
				fmt.Printf("resp Channel Closed. Open Status=%v\n", open)
				respRes = nil
				goto NEXTEXECUTION
			}
			fmt.Printf("response for url=%v is :=%v\n", respRes.inputUrl, len(respRes.responseB))
			fmt.Printf("status is :=%v\n", respRes.status)
			fmt.Printf("executiontime is :=%v\n", respRes.executionTime)
			fmt.Printf("channel open :=%v wg status=%v\n", open, wg)
		case errRes, openerr := <-errChn:
			if !openerr {
				fmt.Printf("resp Channel Closed. Open Status=%v\n", openerr)
				errRes = nil
			} else {
				fmt.Printf("error Occured=%v\n", errRes.errMsg)
				fmt.Printf("error status is :=%v\n", errRes.status)
			}
		}
		counter++
		fmt.Printf("counter=%v\n", counter)
	}
	wg.Wait()
	fmt.Println("Execution Finished")
}

func producerChan(urls []string) (chan string, chan *ErrorWithUrl) {
	urlChn := make(chan string)
	errChn := make(chan *ErrorWithUrl)

	go func() {
		for _, url := range urls {
			urlChn <- url
		}
	}()
	return urlChn, errChn
}

func dataFetcher(producerChn chan string, errChn chan *ErrorWithUrl, awg *sync.WaitGroup) chan *Result {
	fetcherChn := make(chan *Result)
	go func() {
		defer close(producerChn)
		for url := range producerChn {
			awg.Add(1)
			go fetchUrlDataOnChannel(url, fetcherChn, errChn, awg)
		}
	}()
	return fetcherChn
}

func fetchUrlDataOnChannel(url string, fetchChn chan *Result, errChn chan *ErrorWithUrl, awg *sync.WaitGroup) {
	defer awg.Done()
	res, errWithUrl := GetURLResponse(url)
	if errWithUrl != nil {
		errChn <- errWithUrl
		return
	}
	fetchChn <- res
}

type Result struct {
	responseB     string
	status        int
	executionTime time.Duration
	inputUrl      string
}

type ErrorWithUrl struct {
	errMsg error
	status int
}

func GetURLResponse(url string) (*Result, *ErrorWithUrl) {
	startT := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		if resp != nil {
			return nil, &ErrorWithUrl{
				errMsg: err,
				status: resp.StatusCode,
			}
		}
		return nil, &ErrorWithUrl{
			errMsg: err,
			status: -1,
		}
	}
	bodyStr, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, &ErrorWithUrl{
			errMsg: err,
			status: resp.StatusCode,
		}
	}
	res := &Result{
		responseB:     string(bodyStr),
		status:        resp.StatusCode,
		executionTime: time.Now().Sub(startT),
		inputUrl:      url,
	}
	return res, nil
}
