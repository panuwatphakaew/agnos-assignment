package staff

import (
	"context"

	"github.com/panuwatphakaew/agnos-assignment/internal/hospital"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Create(ctx context.Context, staff CreateStaffRequest) (Staff, error)
}

type service struct {
	staffRepository    Repository
	hospitalRepository hospital.Repository
}

func NewService(staffRepo Repository, hospitalRepo hospital.Repository) Service {
	return &service{
		staffRepository:    staffRepo,
		hospitalRepository: hospitalRepo,
	}
}

func (s *service) Create(ctx context.Context, staff CreateStaffRequest) (Staff, error) {
	hospital := hospital.Hospital{
		Name: staff.HospitalName,
	}
	hospitalID, err := s.hospitalRepository.GetOrCreate(ctx, hospital)
	if err != nil {
		return Staff{}, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(staff.Password), bcrypt.DefaultCost)
	if err != nil {
		return Staff{}, err
	}

	staffModel := Staff{
		Username:     staff.Username,
		Password:     string(hashedPassword),
		HospitalName: staff.HospitalName,
		HospitalID:   hospitalID,
	}
	createdStaff, err := s.staffRepository.Create(ctx, staffModel)
	if err != nil {
		return Staff{}, err
	}

	return createdStaff, nil
}
