package auth

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/panuwatphakaew/agnos-assignment/config"
	"github.com/panuwatphakaew/agnos-assignment/internal/staff"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Authenticate(ctx context.Context, loginRequest LoginRequest) (staff.Staff, error)
	GenerateToken(staffID int) (string, error)
}

type service struct {
	staffRepository staff.Repository
	config          config.Config
}

func NewService(staffRepo staff.Repository, config config.Config) Service {
	return &service{
		staffRepository: staffRepo,
		config:          config,
	}
}

func (s *service) Authenticate(ctx context.Context, loginRequest LoginRequest) (staff.Staff, error) {
	res, err := s.staffRepository.GetByUsername(ctx, loginRequest.Username)
	if err != nil {
		return staff.Staff{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(loginRequest.Password))
	if err != nil {
		return staff.Staff{}, err
	}

	return res, nil
}

func (s *service) GenerateToken(staffID int) (string, error) {
	claims := jwt.MapClaims{
		"staff_id": staffID,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.config.JWTSecret))
}
