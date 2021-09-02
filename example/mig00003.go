package main

import (
	gm "github.com/randree/gormigrator"
	"gorm.io/gorm"
)

func init() {

	gm.Mig(gm.State{

		Tag: "users",

		Up: func(db *gorm.DB) error {

			type User struct {
				ID   int `gorm:"primarykey"`
				Name string
			}
			db.AutoMigrate(&User{})

			return db.Create(&User{
				ID:   2009,
				Name: "Peter",
			}).Error
		},

		Down: func(db *gorm.DB) error {
			err := db.Migrator().DropTable("users")
			return err
		},
	})

}
