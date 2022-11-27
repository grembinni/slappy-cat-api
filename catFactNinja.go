package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func ninjaCatFact() CatFact {

	// make request to fact service
	resp, err := http.Get("https://catfact.ninja/fact")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	bodyBytes, _ := io.ReadAll(resp.Body)

	// convert response body to fact
	var catFact CatFact
	json.Unmarshal(bodyBytes, &catFact)
	catFact.ID = uuid.New()
	catFact.Source = "N"
	return catFact
}
