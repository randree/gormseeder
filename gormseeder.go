package gormseeder

import (
	"fmt"
	"os"

	"gorm.io/gorm"
)

const Version = "0.1.0"

var Testing bool = false

func InitSeeder(db *gorm.DB, tablename string) {

	SeederStore := NewSeederStore(db, tablename)

	// If version show
	if getEnv("VERSION", "") == "1" {
		fmt.Println("Gorseeder version: ", Version)
	}

	// If history show
	if getEnv("HISTORY", "") == "1" {
		showHistory(SeederStore)
	}

	tag := getEnv("TAG", "")
	user := getEnv("USER", "")

	if tag == "" {
		fmt.Println("TAG is missing")
		showExample()
		return
	}
	if user == "" {
		fmt.Println("no USER set")
		showExample()
		return
	}

	err := performSeeder(tag, user, db, SeederStore)
	if err != nil {
		fmt.Println("\033[;31mSeeder ERROR! Keep last Seeder step\033[0m")
		fmt.Println(err)
		return
	}
}

func showExample() {
	fmt.Println("Examples:")
	fmt.Println("HISTORY=1 VERSION=1 TAG=[<tag>|all] USER=tester [go run ./...|<build>]")
}

// Get env variable as string
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
