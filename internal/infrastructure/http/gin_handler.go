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

	// User routes
	r.GET("/api/users", userController.GetAllUsers)
	r.POST("/api/users", userController.CreateUser)
	r.GET("/api/users/:id", userController.GetUser)

	// Role routes
	r.POST("/api/roles", roleController.CreateRole)
	r.GET("/api/roles/:id", roleController.GetRole)
	fmt.Println("Routes has been initialized!")
}
