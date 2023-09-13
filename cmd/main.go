package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RafaZeero/go-link-shortener/scripts/utils"
	"github.com/gin-gonic/gin"
)

type Request struct {
	Url string `json:"url"`
}

var links = make(map[string]string)

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
	key := utils.RandSeq(6)

	// Assign to local db
	links[key] = req.Url

	c.JSON(http.StatusOK, gin.H{"newUrl": key})
}
