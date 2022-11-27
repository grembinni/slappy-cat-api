package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitController() {

	router := gin.Default()
	router.GET("cat/facts", getFacts)
	router.GET("cat/facts/:id", getFactByID)
	router.POST("cat/facts", postFacts)
	router.Run("localhost:8080")
}

// responds with the list of all facts as JSON.
func getFacts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, catFacts)
}

// adds a fact from JSON received.
func postFacts(c *gin.Context) {
	var passedFact CatFact

	// Call BindJSON to bind the received JSON to newFact.
	if err := c.BindJSON(&passedFact); err != nil {
		return
	}

	newFact := NewGeneratedCatFact(passedFact.Fact)
	// Add the new fact to the slice.
	catFacts = append(catFacts, newFact)
	c.IndentedJSON(http.StatusCreated, newFact)
}

// locates the fact whose ID value matches the id parameter sent by the client,
// then returns that fact as a response.
func getFactByID(c *gin.Context) {
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
