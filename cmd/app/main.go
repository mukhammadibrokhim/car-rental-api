package main

import (
	"car-rental-api/docs"
	"car-rental-api/internal/infrastructure/database"
	"car-rental-api/internal/infrastructure/http"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

// @title Car Rental API
// @version 1.0
// @description API documentation for the Car Rental service
// @securityDefinitions.bearer BearerAuth
// @in header
// @name Authorization
func main() {

	docs.SwaggerInfo.Title = "Car Rental API"
	docs.SwaggerInfo.Description = "API documentation for the Car Rental service"

	r := gin.Default()
	db := database.SetupDatabase()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	http.InitializeRoutes(r, db)

	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
