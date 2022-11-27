package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// basic cat fact structure
type CatFact struct {
	ID     uuid.UUID `json:"id"`
	Fact   string    `json:"fact"`
	Length float64   `json:"length"`
	Source string    `json:"source"`
}

// cat facts sliced to seed fact database.
var catFacts = []CatFact{
	ninjaCatFact(),
	ninjaCatFact(),
	ninjaCatFact(),
}

func main() {
	router := gin.Default()
	router.GET("cat/facts", getFacts)
	router.GET("cat/facts/:id", getAlbumByID)
	router.POST("cat/facts", postFacts)
	router.Run("localhost:8080")
}

// responds with the list of all facts as JSON.
func getFacts(c *gin.Context) {
	newFact := ninjaCatFact()
	catFacts = append(catFacts, newFact)
	c.IndentedJSON(http.StatusOK, catFacts)
}

// adds a fact from JSON received.
func postFacts(c *gin.Context) {
	var newFact CatFact

	// Call BindJSON to bind the received JSON to newFact.
	if err := c.BindJSON(&newFact); err != nil {
		return
	}

	// Add the new fact to the slice.
	catFacts = append(catFacts, newFact)
	c.IndentedJSON(http.StatusCreated, newFact)
}

// locates the fact whose ID value matches the id parameter sent by the client,
// then returns that fact as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	// Loop over the list of facts, looking for
	// an fact whose ID value matches the parameter.
	for _, a := range catFacts {
		if a.ID.String() == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "fact not found"})
}
