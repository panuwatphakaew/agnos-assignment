package patient

import (
	"context"

	"github.com/uptrace/bun"
)

type Repository interface {
	Get(ctx context.Context, input Patient) ([]Patient, error)
}

type repository struct {
	db *bun.DB
}

func NewRepository(db *bun.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Get(ctx context.Context, input Patient) ([]Patient, error) {
	var patients []Patient
	err := r.db.NewSelect().Model(&patients).Where("national_id = ? OR passport_id = ? OR first_name = ? OR middle_name = ? OR last_name = ? OR date_of_birth = ? OR phone_number = ? OR email = ?", input.NationalID, input.PassportID, input.FirstName, input.MiddleName, input.LastName, input.DateOfBirth, input.PhoneNumber, input.Email).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return patients, nil
}
