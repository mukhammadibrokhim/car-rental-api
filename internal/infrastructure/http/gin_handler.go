package http

import (
	"car-rental-api/internal/interfaces"
	"car-rental-api/internal/repository"
	"car-rental-api/internal/usecase"
	"fmt"
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func InitializeRoutes(r *gin.Engine, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userController := interfaces.NewUserController(userUsecase)

	roleRepo := repository.NewRoleRepository(db)
	roleUsecase := usecase.NewRoleUsecase(roleRepo)
	roleController := interfaces.NewRoleController(roleUsecase)

	authUsecase := usecase.NewAuthUsecase(userRepo)
	authController := interfaces.NewAuthController(authUsecase)

	api := r.Group("/api")
	{
		authGroup := api.Group("/auth")
		{
			authGroup.POST("/login", authController.Login)
			authGroup.POST("/register", authController.Register)
		}

		users := api.Group("/users", AuthMiddleware())
		{
			users.GET("", userController.GetAllUsers)
			users.POST("", userController.CreateUser)
			users.GET(":id", userController.GetUser)
		}

		roles := api.Group("/roles", AuthMiddleware())
		{
			roles.POST("", roleController.CreateRole)
			roles.GET(":id", roleController.GetRoleId)
		}

		bookings := api.Group("/bookings", AuthMiddleware())
		{
			bookings.GET("")
			bookings.POST("")
			bookings.GET(":id")
		}
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"error": "Route not found"})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	fmt.Println("Routes have been initialized!")
}
