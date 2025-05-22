package staff

import (
	"context"

	"github.com/uptrace/bun"
)

type Repository interface {
	Create(ctx context.Context, staff Staff) (Staff, error)
	GetByUsername(ctx context.Context, username string) (Staff, error)
}

type repository struct {
	db *bun.DB
}

func NewRepository(db *bun.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(ctx context.Context, staff Staff) (Staff, error) {
	_, err := r.db.NewInsert().Model(&staff).Returning("*").Exec(ctx)
	if err != nil {
		return Staff{}, err
	}
	return staff, nil
}

func (r *repository) GetByUsername(ctx context.Context, username string) (Staff, error) {
	var staff Staff
	err := r.db.NewSelect().Model(&staff).Where("username = ?", username).Scan(ctx)
	if err != nil {
		return Staff{}, err
	}
	return staff, nil
}
