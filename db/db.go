package db

import (
	"database/sql"

	"github.com/panuwatphakaew/agnos-assignment/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func Connect(cfg config.Config) *bun.DB {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(cfg.DbURL)))

	db := bun.NewDB(sqldb, pgdialect.New())
	return db
}
