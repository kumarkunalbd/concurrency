package network

import "testing"

func TestMultiURLTime(t *testing.T) {
	inputUrls := []string{
		"https://catfact.ninja/fact",
		"https://api.coindesk.com/v1/bpi/currentprice.json",
		"https://api.publicapis.org/entries",
		"https://api.agify.io/?name=meelad",
		"https://api.genderize.io/?name=luc",
		"https://api.nationalize.io/?name=nathaniel",
		"https://datausa.io/api/data?drilldowns=Nation&measures=Population",
	}
	MultiURLTime(inputUrls)
}
