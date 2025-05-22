package router

import (
	"github.com/gin-gonic/gin"
	"github.com/panuwatphakaew/agnos-assignment/config"
	"github.com/panuwatphakaew/agnos-assignment/internal/auth"
	"github.com/panuwatphakaew/agnos-assignment/internal/patient"
	"github.com/panuwatphakaew/agnos-assignment/internal/staff"
	"github.com/uptrace/bun"
)

func Setup(db *bun.DB, config config.Config) *gin.Engine {
	r := gin.Default()
	r.Use(JWTAuthMiddleware(config))

	staffGroup := r.Group("/staff")
	patientGroup := r.Group("/patient")

	auth.RegisterRoutes(staffGroup, db, config)

	staff.RegisterRoutes(staffGroup, db)
	patient.RegisterRoutes(patientGroup, db)

	return r
}
