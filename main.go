package main

// cat facts sliced to seed fact database.
var catFacts = []CatFact{}

func main() {
	catFacts = GetNinjaCatFacts()
	catFacts = append(catFacts, GetNinjaCatFact())
	InitController()
}
