package gormseeder

import (
	"errors"
	"fmt"
	"strings"
)

func showHistory(store *SeederStore) error {

	list, _ := store.FetchAll()
	if len(list) == 0 {
		return errors.New("there is no seed done yet")
	}
	fmt.Printf("\n| %-40s | %-40s | %-20s |\n", "DATETIME HISTORY", "FILENAME", "USER")
	fmt.Printf("| %-40s | %-40s | %-20s |\n", strings.Repeat("-", 40), strings.Repeat("-", 40), strings.Repeat("-", 20))
	for _, entry := range list {
		fmt.Printf("| %-40s | %-30s           | %-20s |\n", entry.CreatedAt, entry.Filename, entry.User)
	}
	return nil
}
