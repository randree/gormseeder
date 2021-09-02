package main

import (
	gs "github.com/randree/gormseeder"
	"gorm.io/gorm"
)

func init() {

	gs.Seed(gs.State{

		Tag: "mock_products",

		Perform: func(db *gorm.DB) error {
			type Product struct {
				ID   uint   `gorm:"primarykey"`
				Name string `gorm:"size:255"`
			}
			err := db.AutoMigrate(&Product{})
			if err != nil {
				return err
			}

			var product = []Product{{Name: "testprod1"}, {Name: "testprod2"}, {Name: "testprod3"}}
			db.Create(&product)

			return err
		},
	})

}
