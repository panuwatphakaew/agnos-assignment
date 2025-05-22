package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/panuwatphakaew/agnos-assignment/config"
	"github.com/panuwatphakaew/agnos-assignment/internal/staff"
	"github.com/uptrace/bun"
)

func RegisterRoutes(r *gin.RouterGroup, db *bun.DB, config config.Config) {
	staffRepository := staff.NewRepository(db)
	service := NewService(staffRepository, config)
	handler := NewHandler(service)

	r.POST("/login", handler.Login)
}

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) Login(c *gin.Context) {
	var loginRequest LoginRequest
	err := c.ShouldBindJSON(&loginRequest)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	staff, err := h.service.Authenticate(c.Request.Context(), loginRequest)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := h.service.GenerateToken(staff.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
