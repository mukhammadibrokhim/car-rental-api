package interfaces

import (
	"car-rental-api/internal/domain"
	"car-rental-api/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type RoleController struct {
	RoleUsecase usecase.RoleUsecase
}

func NewRoleController(uc usecase.RoleUsecase) *RoleController {
	return &RoleController{RoleUsecase: uc}
}

func (c *RoleController) CreateRole(ctx *gin.Context) {
	var role domain.Role

	if err := ctx.ShouldBindJSON(&role); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.RoleUsecase.CreateRole(&role); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create role"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Role created successfully"})
}

func (c *RoleController) GetRole(ctx *gin.Context) {
	id := ctx.Param("id")

	uintId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	role, err := c.RoleUsecase.GetRoleByID(uint(uintId))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}
	ctx.JSON(http.StatusOK, role)
}
