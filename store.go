package gormseeder

import (
	"gorm.io/gorm"
)

// Seeder model
type Seeder struct {
	gorm.Model
	Tag      string `gorm:"size:455"`
	Filename string `gorm:"size:455"`
	User     string `gorm:"size:255"`
}

type SeederStore struct {
	tableName string
	db        *gorm.DB
}

func NewSeederStore(db *gorm.DB, tableName string) *SeederStore {

	if tableName == "" {
		tableName = "Seeders"
	}

	// Creating a Seeder table if not exists
	db.Table(tableName).AutoMigrate(&Seeder{})

	return &SeederStore{
		tableName: tableName,
		db:        db,
	}

}

// If tag is not found an NoRecordFound
func (m *SeederStore) TagInDB(tag string) bool {
	seed := &Seeder{}
	if err := m.db.Table(m.tableName).Take(&seed, "tag = ?", tag).Error; err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

func (m *SeederStore) FetchAll() ([]*Seeder, error) {
	var SeederList []*Seeder
	if err := m.db.Table(m.tableName).Order("id DESC").Find(&SeederList).Error; err != nil {
		return nil, err
	}
	return SeederList, nil
}

func (m *SeederStore) SaveState(tag, filename, user string) error {
	currentState := &Seeder{
		Tag:      tag,
		Filename: filename,
		User:     user,
	}
	if err := m.db.Table(m.tableName).Create(&currentState).Error; err != nil {
		return err
	}
	return nil
}
