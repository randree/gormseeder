package main

import (
	gm "github.com/randree/gormigrator"
	"gorm.io/gorm"
)

func init() {

	gm.Mig(gm.State{

		Tag: "customers",

		Up: func(db *gorm.DB) error {

			type Customer struct {
				ID   int `gorm:"primarykey"`
				Name string
			}

			db.AutoMigrate(&Customer{})

			return db.Create(&Customer{
				ID:   2009,
				Name: "Cust1",
			}).Error
		},

		Down: func(db *gorm.DB) error {
			err := db.Migrator().DropTable("customers")

			return err
		},
	})

}
