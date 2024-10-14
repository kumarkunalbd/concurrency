package main

import (
	"ConcurrencyP/pipeline"
	"fmt"
	"log"
)

func main() {
	generateChn := pipeline.GenerateData()
	prepData := pipeline.PrepareData(generateChn)
	disData := pipeline.DisplayData(prepData)

	for data := range disData {
		log.Printf("Items: %+v", data)
	}

	fmt.Println("execution Finished")
}
