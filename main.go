package main

import (
	"github.com/panuwatphakaew/agnos-assignment/config"
	"github.com/panuwatphakaew/agnos-assignment/db"
	"github.com/panuwatphakaew/agnos-assignment/router"
)

func main() {
	cfg := config.LoadConfig()
	db := db.Connect(cfg)

	r := router.Setup(db, cfg)
	r.Run(":" + cfg.Port)
}
