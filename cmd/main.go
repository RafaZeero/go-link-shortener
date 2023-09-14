package main

import (
	"github.com/RafaZeero/go-link-shortener/internal/database"
	httpapi "github.com/RafaZeero/go-link-shortener/internal/http-api"
)

func main() {
	httpapi.InitAPI()

	if err := database.Connect(); err != nil {
		panic(err)
	}
	defer database.Close()

}
