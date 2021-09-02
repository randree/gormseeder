package gormseeder

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strings"

	"gorm.io/gorm"
)

type State struct {
	Tag      string
	Filename string
	Perform  func(*gorm.DB) error
}

var Seeders = []State{}

func Seed(state State) {

	// 1. Check Tag
	filename := caller()
	if state.Tag == "" {
		log.Fatal("tag is missing (" + filename + ")")
	}
	if state.Tag == "all" {
		log.Fatal("tag can't have reserved name all (" + filename + ")")
	}
	if strings.Contains(state.Tag, " ") {
		log.Fatal("tag contains a whitespace (" + filename + ")")
	}
	if tagExists(state.Tag) {
		log.Fatal("tag already exists (" + filename + ")")
	}

	// 2. Check Perfom
	if state.Perform == nil {
		log.Fatal("up function missing (" + filename + ")")
	}

	state.Filename = filename

	Seeders = append(Seeders, state)
}

func tagExists(newTag string) bool {
	for _, Seeder := range Seeders {
		if Seeder.Tag == newTag {
			return true
		}
	}
	return false
}

func caller() string {
	_, path, _, _ := runtime.Caller(2)
	return filepath.Base(path)
}

func performSeeder(tag, user string, db *gorm.DB, SeederStore *SeederStore) error {

	// No username
	if user == "" {
		return fmt.Errorf("no username set")
	}

	// If there is no Seeder yet current tag will be empty thus "null"
	if tag == "" {
		return fmt.Errorf("no TAG defined")
	}

	if !tagExists(tag) && tag != "all" {
		return fmt.Errorf("can't find TAG in files")
	}

	for i, seeder := range Seeders {
		if seeder.Tag == tag || tag == "all" {
			// Check if tag is in DB
			if SeederStore.TagInDB(seeder.Tag) {
				fmt.Println("\033[;32mo seed with TAG (", seeder.Tag, ") already performed\033[0m")
				continue
			}
			fmt.Println("\033[;33m+ Perform seed with TAG:", seeder.Tag, "\033[0m")
			err := Seeders[i].Perform(db)
			if err != nil {
				return err
			}
			SeederStore.SaveState(seeder.Tag, seeder.Filename, user)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
