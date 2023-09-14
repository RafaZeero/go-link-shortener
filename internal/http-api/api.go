package httpapi

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitAPI() {
	// Init router
	r := gin.Default()

	routes(r)

	log.Fatal(r.Run(":3000"))
}

func routes(r *gin.Engine) {
	rs := r.Group("")
	{
		rs.POST("/", CreateNewUrl)
		rs.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "pong")
		})
		rs.GET("/:code", RedirectToUrl)
	}
}
