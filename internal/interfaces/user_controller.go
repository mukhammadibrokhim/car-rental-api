package interfaces

import (
	"car-rental-api/internal/domain"
	"car-rental-api/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController struct {
	UserUsecase usecase.UserUsecase
}

func NewUserController(uc usecase.UserUsecase) *UserController {
	return &UserController{UserUsecase: uc}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user domain.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.UserUsecase.CreateUser(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func (c *UserController) GetUser(ctx *gin.Context) {

	id := ctx.Param("id")

	uintId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.UserUsecase.GetUserByID(uint(uintId))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// @Summary Get all users
// @Description Get a list of users
// @ID get-users
// @Accept  json
// @Produce  json
// @Success 200 {array} User // Replace User with your actual response type
// @Router /api/users [get]
func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.UserUsecase.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get all users"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": users})
}
