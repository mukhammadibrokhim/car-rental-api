package main

import (
	"car-rental-api/internal/infrastructure/database"
	"car-rental-api/internal/infrastructure/http"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	db := database.SetupDatabase()

	http.InitializeRoutes(r, db)

	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
