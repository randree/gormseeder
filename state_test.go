package gormseeder

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func createMockSeederList() {

	type migtest struct {
		ID    uint   `gorm:"primarykey,autoIncrement"`
		Name  string `gorm:"size:255"`
		Email string `gorm:"size:300"`
	}

	Seed(State{
		Filename: "seed001",
		Tag:      "create_migtest",
		Perform: func(db *gorm.DB) error {
			err := db.AutoMigrate(&migtest{})
			return err
		},
	})
	Seed(State{
		Filename: "seed002",
		Tag:      "add_entries",
		Perform: func(db *gorm.DB) error {
			err := db.Create(&migtest{Name: "Tester", Email: "test@test.tx"}).Error
			return err
		},
	})
	Seed(State{
		Filename: "seed003",
		Tag:      "add_column",
		Perform: func(db *gorm.DB) error {
			type migtest struct {
				Testcol string `gorm:"size:200"`
			}
			err := db.Migrator().AddColumn(&migtest{}, "testcol")
			return err
		},
	})
	Seed(State{
		Filename: "seed004",
		Tag:      "add_entries_2",
		Perform: func(db *gorm.DB) error {
			err := db.Create(&migtest{Name: "Tester2", Email: "test@test.tx2"}).Error
			return err
		},
	})
}

func Test_performSeeder(t *testing.T) {

	createMockSeederList()

	db, err := gorm.Open(postgres.Open("host=localhost user=user password=passpass dbname=testdb port=5432 sslmode=disable"), &gorm.Config{Logger: logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			LogLevel: logger.Silent, // Log level
			// Colorful: true,
			// LogLevel: logger.Info,
		},
	)})
	t.Run("DB ok", func(t *testing.T) {
		assert.NoError(t, err)
	})

	SeederTableName := "seed-test-table"

	// Pre clean up
	err = db.Migrator().DropTable(SeederTableName)
	t.Run("Migrator 1", func(t *testing.T) {
		assert.NoError(t, err)
	})
	db.Migrator().DropTable("migtests")
	t.Run("Migrator 2", func(t *testing.T) {
		assert.NoError(t, err)
	})

	SeederStore := NewSeederStore(db, SeederTableName)

	type args struct {
		tag  string
		user string
	}
	tests := []struct {
		name       string
		args       args
		currentTag string // Current tag AFTER! Seeder is executed
		wantErr    bool
	}{
		{
			name:    "empty tag",
			wantErr: true,
			args: args{
				tag:  "",
				user: "admin",
			},
			currentTag: "",
		},
		{
			name:    "Add entries with fail",
			wantErr: true, // cant add entries to non existent table
			args: args{
				tag:  "add_entries",
				user: "admin",
			},
			currentTag: "",
		},
		{
			name:    "create a table",
			wantErr: false,
			args: args{
				tag:  "create_migtest",
				user: "admin",
			},
			currentTag: "",
		},
		{
			name:    "Add entries",
			wantErr: false,
			args: args{
				tag:  "add_entries",
				user: "admin",
			},
			currentTag: "",
		},
		{
			name:    "Repeat add entries",
			wantErr: false,
			args: args{
				tag:  "add_entries",
				user: "admin",
			},
			currentTag: "",
		},
		{
			name:    "create a table again",
			wantErr: false,
			args: args{
				tag:  "create_migtest",
				user: "admin",
			},
			currentTag: "",
		},
		{
			name:    "non existing tag",
			wantErr: true,
			args: args{
				tag:  "fooBar",
				user: "admin",
			},
			currentTag: "",
		},
		{
			name:    "perform all",
			wantErr: false,
			args: args{
				tag:  "all",
				user: "admin",
			},
			currentTag: "",
		},
		{
			name:    "perform all repeat",
			wantErr: false,
			args: args{
				tag:  "all",
				user: "admin",
			},
			currentTag: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := performSeeder(tt.args.tag, tt.args.user, db, SeederStore)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

		})
	}

	// After clean up
	err = db.Migrator().DropTable(SeederTableName)
	t.Run("Migrator 1", func(t *testing.T) {
		assert.NoError(t, err)
	})
	db.Migrator().DropTable("migtest")
	t.Run("Migrator 2", func(t *testing.T) {
		assert.NoError(t, err)
	})
}
