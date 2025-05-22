package hospital

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
)

type Repository interface {
	GetOrCreate(ctx context.Context, hospital Hospital) (int, error)
}

type repository struct {
	db *bun.DB
}

func NewRepository(db *bun.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetOrCreate(ctx context.Context, hospital Hospital) (int, error) {
	var id int
	err := r.db.NewSelect().Model(&hospital).Where("name = ?", hospital.Name).Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			_, err := r.db.NewInsert().Model(&hospital).Returning("id").Exec(ctx)
			if err != nil {
				return 0, err
			}
			id = hospital.ID
		} else {
			return 0, err
		}
	}
	return id, nil
}
