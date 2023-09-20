package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type detail struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	City   string `json:"city"`
	Age    int    `json:"age"`
	GitHub string `json:"github"`
}

var details = []detail{
	{ID: "1", Name: "test1", City: "pune", Age: 20, GitHub: "test1"},
	{ID: "2", Name: "test2", City: "pune", Age: 20, GitHub: "test2"},
}

func getDetails(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, details)
}

func addDetails(c *gin.Context) {
	var newDetail detail

	if err := c.BindJSON(&newDetail); err != nil {
		return
	}

	details = append(details, newDetail)
	c.IndentedJSON(http.StatusCreated, newDetail)
}

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Endpoint hit")
// }

// func handleRequest() {
// 	http.HandleFunc("/", homePage)
// 	log.Fatal(http.ListenAndServe(":8000", nil))
// }

func main() {
	router := gin.Default()
	router.GET("/info", getDetails)
	router.POST("/info", addDetails)
	router.Run("localhost:8080")
}
