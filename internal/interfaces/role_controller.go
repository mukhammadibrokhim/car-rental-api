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

// CreateRole godoc
// @Summary Create a new Role
// @Description Create a role for user permissions
// @Tags Role
// @Accept json
// @Produce json
// @Param role body domain.Role true "Role data" // Assuming domain.Role is the structure representing a Role
// @Success 201 {object} domain.Role "Created role details"
// @Failure 400 {object} payload.ErrorResponse "Invalid request"
// @Failure 500 {object} payload.ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /api/roles [post]
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

// GetRoleId godoc
// @Summary Get Role by ID
// @Description Get role details using its ID
// @Tags Role
// @Accept json
// @Produce json
// @Param id path int true "Role ID"
// @Success 200 {object} domain.Role "Role details"  // Assuming domain.Role is the role struct you're returning
// @Failure 400 {object} payload.ErrorResponse "Invalid ID"
// @Failure 404 {object} payload.ErrorResponse "Role not found"
// @Security ApiKeyAuth
// @Router /api/roles/{id} [get]
func (c *RoleController) GetRoleId(ctx *gin.Context) {
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
