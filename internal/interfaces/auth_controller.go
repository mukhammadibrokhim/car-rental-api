package interfaces

import (
	"car-rental-api/internal/domain/payload"
	"car-rental-api/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
	AuthUsecase usecase.AuthUsecase
}

func NewAuthController(authUsecase usecase.AuthUsecase) *AuthController {
	return &AuthController{AuthUsecase: authUsecase}
}

// Login godoc
// @Summary Login with credentials
// @Description Authenticate the user and return a token
// @Tags Auth
// @Accept json
// @Produce json
// @Param credentials body payload.LoginRequest true "Login credentials"
// @Success 200 {object} payload.LoginResponse
// @Failure 400 {object} payload.ErrorResponse
// @Router /api/auth/login [post]
func (u *AuthController) Login(c *gin.Context) {
	var request payload.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body!"})
		return
	}
	token, err := u.AuthUsecase.Login(request.Email, request.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}

// Register godoc
// @Summary Register with credentials
// @Description Register the user and return a token
// @Tags Auth
// @Accept json
// @Produce json
// @Param credentials body payload.RegisterRequest true "Register "
// @Success 200 {object} payload.RegisterResponse
// @Failure 400 {object} payload.ErrorResponse
// @Router /api/auth/register [post]
func (u *AuthController) Register(c *gin.Context) {
	var request payload.RegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body!"})
		return
	}
	token, err := u.AuthUsecase.Register(request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registered Successfully!", "token": token})
}
