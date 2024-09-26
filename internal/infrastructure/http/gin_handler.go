package http

import (
	_ "car-rental-api/cmd/app/docs" // Ensure this is correctly imported for swagger
	"car-rental-api/internal/interfaces"
	"car-rental-api/internal/repository"
	"car-rental-api/internal/usecase"
	"fmt"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
			roles.GET(":id", roleController.GetRole)
		}
		bookings := api.Group("/bookings", AuthMiddleware())
		{
			bookings.GET("")
			bookings.POST("")
			bookings.GET(":id")
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"error": "Route not found"})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	fmt.Println("Routes has been initialized!")
}
