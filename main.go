package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// basic cat fact structure
type catFact struct {
	ID       string `json:"id"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

// cat facts sliced to seed fact database.
var catFacts = []catFact{
	{ID: "1", Question: "Blue Train", Answer: "John Coltrane"},
	{ID: "2", Question: "Jeru", Answer: "Gerry Mulligan"},
	{ID: "3", Question: "Sarah Vaughan and Clifford Brown", Answer: "Sarah Vaughan"},
}

func main() {
	router := gin.Default()
	router.GET("cat/facts", getFacts)
	router.POST("cat/facts", postFacts)
	router.Run("localhost:8080")
}

// responds with the list of all facts as JSON.
func getFacts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, catFacts)
}

// postAlbums adds an album from JSON received in the request body.
func postFacts(c *gin.Context) {
	var newFact catFact

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newFact); err != nil {
		return
	}

	// Add the new album to the slice.
	catFacts = append(catFacts, newFact)
	c.IndentedJSON(http.StatusCreated, newFact)
}
