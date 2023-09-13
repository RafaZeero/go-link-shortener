package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Url string `json:"url"`
}

var links = make(map[string]string)
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func main() {
	r := gin.Default()

	r.POST("/", CreateNewUrl)
	r.GET("/:code", RedirectToUrl)

	log.Fatal(r.Run(":3000"))
}

func RedirectToUrl(c *gin.Context) {
	code := c.Param("code")

	url, found := links[code]
	if !found {
		fmt.Println("Link not found")
		c.JSON(http.StatusNotFound, gin.H{"Error": "Link not found"})
		return
	}

	c.Redirect(http.StatusFound, url)
}

func CreateNewUrl(c *gin.Context) {
	var req Request

	if err := c.BindJSON(&req); err != nil {
		fmt.Println("Error binding data", err.Error())
		c.JSON(http.StatusOK, gin.H{"Error": err.Error()})
		return
	}

	// Generate random string
	key := randSeq(6)

	// Assign to local db
	links[key] = req.Url

	c.JSON(http.StatusOK, gin.H{"newUrl": key})
}

// Create random string sequence
func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
