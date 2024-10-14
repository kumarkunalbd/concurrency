package network

import (
	"fmt"
	"sync"
)

func MultiUrlFetch(urls []string) {
	var wg sync.WaitGroup
	producedChn, errChn := producerChanWithWG(urls, &wg)
	fetcherChn := dataFetchWithPubSub(producedChn, errChn)
	processFetchedData(fetcherChn, errChn, &wg)
	wg.Wait()
	fmt.Println("Execution finised")
}

func producerChanWithWG(urls []string, awg *sync.WaitGroup) (chan string, chan *ErrorWithUrl) {
	urlChn := make(chan string)
	errChn := make(chan *ErrorWithUrl)
	awg.Add(len(urls))
	go func() {
		for _, url := range urls {
			urlChn <- url
		}
	}()
	return urlChn, errChn
}

func dataFetchWithPubSub(prdChn chan string, errChn chan *ErrorWithUrl) chan *Result {
	datfetchChn := make(chan *Result)
	go func() {
		defer close(prdChn)
		for url := range prdChn {
			go fetchDataOnChannelPubSub(url, datfetchChn, errChn)
		}
	}()
	return datfetchChn
}

func fetchDataOnChannelPubSub(url string, fetchChn chan *Result, errChn chan *ErrorWithUrl) {
	res, anErr := GetURLResponse(url)
	if anErr != nil {
		errChn <- anErr
		return
	}
	fetchChn <- res
}

func processFetchedData(producerChan chan *Result, errChn chan *ErrorWithUrl, awg *sync.WaitGroup) {
	go func() {
		defer close(producerChan)
		for res := range producerChan {
			fmt.Printf("response for url=%v status=%v executiontime=%v and len of bidy=%v\n", res.inputUrl, res.status, res.executionTime, len(res.responseB))
			awg.Done()
		}
	}()
	go func() {
		defer close(errChn)
		for errM := range errChn {
			fmt.Printf("err for status=%v err=%v \n", errM.status, errM.errMsg)
			awg.Done()
		}
	}()
}

func processError(errorChan chan *ErrorWithUrl, awg *sync.WaitGroup) {
	go func() {
		defer close(errorChan)
		for errM := range errorChan {
			fmt.Printf("err for status=%v err=%v \n", errM.status, errM.errMsg)
		}
	}()
	awg.Wait()
}
