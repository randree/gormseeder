package main

import (
	gs "github.com/randree/gormseeder"
	"gorm.io/gorm"
)

func init() {

	type migtest struct {
		ID    uint   `gorm:"primarykey,autoIncrement"`
		Name  string `gorm:"size:255"`
		Email string `gorm:"size:300"`
	}

	gs.Seed(gs.State{

		Tag: "create_migtest",

		Perform: func(db *gorm.DB) error {
			err := db.AutoMigrate(&migtest{})
			return err
		},
	})

}
