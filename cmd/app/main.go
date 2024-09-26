package main

import (
	_ "car-rental-api/cmd/app/docs" // Import the generated docs
	"car-rental-api/internal/infrastructure/database"
	"car-rental-api/internal/infrastructure/http"
	"github.com/gin-gonic/gin"
	"log"
)

// @title Car Rental API
// @version 1.0
// @description API documentation for the Car Rental service
// @contact.name API Support
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /api
// @schemes http
func main() {
	r := gin.Default()
	db := database.SetupDatabase()

	http.InitializeRoutes(r, db)

	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
