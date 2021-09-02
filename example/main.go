package main

import (
	"fmt"
	"log"
	"os"

	gs "github.com/randree/gormseeder"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	db, err := gorm.Open(postgres.Open("host=localhost user=user password=passpass dbname=testdb port=5432 sslmode=disable"), &gorm.Config{Logger: logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{LogLevel: logger.Silent},
	)})
	if err != nil {
		fmt.Println(err.Error())
	}

	gs.InitSeeder(db, "Seeders")
}
