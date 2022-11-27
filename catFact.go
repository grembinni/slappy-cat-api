package main

import "github.com/google/uuid"

// basic cat fact structure
type CatFact struct {
	ID     uuid.UUID `json:"id"`
	Fact   string    `json:"fact"`
	Length int       `json:"length"`
	Source string    `json:"source"`
}

// constructor for ninja fact
func NewNinjaCatFact(fact string) CatFact {
	return newCatFact(fact, "N")
}

// constructor for ninja fact
func NewGeneratedCatFact(fact string) CatFact {
	return newCatFact(fact, "G")
}

// constructor for ninja fact
func newCatFact(fact string, source string) CatFact {
	var catFact CatFact
	catFact.ID = uuid.New()
	catFact.Fact = fact
	catFact.Length = len(fact)
	catFact.Source = source
	return catFact
}
