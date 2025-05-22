package staff

import (
	"github.com/gin-gonic/gin"
	"github.com/panuwatphakaew/agnos-assignment/internal/hospital"
	"github.com/uptrace/bun"
)

func RegisterRoutes(r *gin.RouterGroup, db *bun.DB) {
	staffRepository := NewRepository(db)
	hospitalRepository := hospital.NewRepository(db)
	service := NewService(staffRepository, hospitalRepository)
	handler := NewHandler(service)

	r.POST("/create", handler.Create)
}

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) Create(c *gin.Context) {
	var staff CreateStaffRequest
	err := c.ShouldBindJSON(&staff)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	res, err := h.service.Create(c, staff)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create staff"})
		return
	}

	c.JSON(200, gin.H{"id": res.ID})
}

type CreateStaffRequest struct {
	Username     string `json:"username" binding:"required"`
	Password     string `json:"password" binding:"required"`
	HospitalName string `json:"hospital_name" binding:"required"`
}
