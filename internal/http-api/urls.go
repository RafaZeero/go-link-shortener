package httpapi

import (
	"fmt"
	"net/http"

	"github.com/RafaZeero/go-link-shortener/internal/database"
	"github.com/RafaZeero/go-link-shortener/scripts/utils"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

type Request struct {
	Url string `json:"url"`
}

func RedirectToUrl(c *gin.Context) {
	code := c.Param("code")

	var url string

	tx, _ := database.DB.Begin()

	q := `SELECT link FROM links WHERE short = $1`

	stmt, _ := tx.Prepare(q)

	row, err := stmt.Query(code)
	if err != nil {
		fmt.Println("Error getting data from DB", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"Error in query": err.Error()})
		return
	}

	for row.Next() {
		err := row.Scan(&url)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error in scan": err.Error()})
			return
		}
	}

	tx.Commit()

	c.Redirect(http.StatusFound, url)
}

func CreateNewUrl(c *gin.Context) {
	var req Request

	tx, _ := database.DB.Begin()

	if err := c.BindJSON(&req); err != nil {
		fmt.Println("Error binding data", err.Error())
		c.JSON(http.StatusOK, gin.H{"Error": err.Error()})
		return
	}

	// Generate random string
	code := utils.RandSeq(6)

	q := `INSERT INTO links (links, short) VALUES ($1, $2)`

	stmt, _ := tx.Prepare(q)

	// Assign to local db
	row, err := stmt.Exec(req.Url, code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx.Commit()

	id, _ := row.LastInsertId()

	c.JSON(http.StatusCreated, gin.H{"newUrl": code, "id": id})
}
