package main

import (
	"ConcurrencyP/hashing"
	"fmt"
)

func main() {
	inputUrls := []string{
		"https://catfact.ninja/fact",
		"https://api.coindesk.com/v1/bpi/currentprice.json",
		"https://api.publicapis.org/entries",
		"https://api.agify.io/?name=meelad",
		"https://api.genderize.io/?name=luc",
		"https://api.nationalize.io/?name=nathaniel",
		"https://datausa.io/api/data?drilldowns=Nation&measures=Population",
		"http://halua.com",
		"http://http://aaaaaaaaallllllll00.com/",
	}

	for _, url := range inputUrls {
		res := hashing.PerformHash(url)
		fmt.Printf("url =%v  hashed=%v\n", url, res)
	}
}
