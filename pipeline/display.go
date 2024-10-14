package pipeline

import (
	"fmt"
	"sync"
)

// tp access in different files
func DisplayData(ic <-chan Operation) <-chan string {
	oc := make(chan string)

	go func() {
		wg := &sync.WaitGroup{}
		for input := range ic {
			wg.Add(1)
			go concatenateValue(input, oc, wg)
		}
		wg.Wait()
		close(oc)
	}()
	return oc
}

func concatenateValue(input Operation, output chan<- string, wg *sync.WaitGroup) {
	concatenated := fmt.Sprintf("id: %d, multiplication: %d, addition %d", input.id, input.multiply, input.addition)
	output <- concatenated
	wg.Done()
}
