package main

import (
	gs "github.com/randree/gormseeder"
	"gorm.io/gorm"
)

func init() {

	gs.Seed(gs.State{

		Tag: "mock_users",

		Perform: func(db *gorm.DB) error {
			type User struct {
				ID   uint   `gorm:"primarykey"`
				Name string `gorm:"size:255"`
			}
			err := db.AutoMigrate(&User{})
			if err != nil {
				return err
			}

			var users = []User{{Name: "testuser1"}, {Name: "testuser2"}, {Name: "testuser3"}}
			db.Create(&users)

			return err
		},
	})

}
