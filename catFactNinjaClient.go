package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// basic cat fact structure
type NinjaCatFact struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

// basic cat fact structure
type NinjaCatFactResponse struct {
	Data []NinjaCatFact `json:"data"`
}

func GetNinjaCatFact() CatFact {

	// make request to fact service
	resp, err := http.Get("https://catfact.ninja/fact")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	bodyBytes, _ := io.ReadAll(resp.Body)

	// convert response body to fact
	var catFact NinjaCatFact
	json.Unmarshal(bodyBytes, &catFact)
	return NewNinjaCatFact(catFact.Fact)
}

func GetNinjaCatFacts() []CatFact {

	// make request to fact service
	resp, err := http.Get("https://catfact.ninja/facts")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	bodyBytes, _ := io.ReadAll(resp.Body)

	fmt.Println("response: ", string(bodyBytes))

	// convert response body to fact
	var passedCatFacts NinjaCatFactResponse
	json.Unmarshal(bodyBytes, &catFacts)

	newCatFacts := make([]CatFact, len(passedCatFacts.Data))
	for _, passedFact := range passedCatFacts.Data {

		fmt.Println("passed fact: ", passedFact)
		newCatFacts = append(newCatFacts, NewNinjaCatFact(passedFact.Fact))
	}
	return newCatFacts
}
