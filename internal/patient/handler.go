package patient

import (
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func RegisterRoutes(r *gin.RouterGroup, db *bun.DB) {
	repository := NewRepository(db)
	service := NewService(repository)
	handler := NewHandler(service)

	r.GET("/patients/search", handler.GetPatients)
}

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) GetPatients(c *gin.Context) {
	var req GetPatientsRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	input := Patient{
		NationalID:  req.NationalID,
		PassportID:  req.PassportID,
		FirstName:   req.FirstName,
		MiddleName:  req.MiddleName,
		LastName:    req.LastName,
		DateOfBirth: req.DateOfBirth,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
	}

	patients, err := h.service.GetPatients(c, input)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get patients"})
		return
	}

	c.JSON(200, patients)
}

type GetPatientsRequest struct {
	NationalID  string `form:"nationalId"`
	PassportID  string `form:"passportId"`
	FirstName   string `form:"firstName"`
	MiddleName  string `form:"middleName"`
	LastName    string `form:"lastName"`
	DateOfBirth string `form:"dateOfBirth"`
	PhoneNumber string `form:"phoneNumber"`
	Email       string `form:"email"`
}
