package crud

import (
	"github.com/JeremyPersing/gographqltut/graph/model"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func InsertDog(db *gorm.DB, name string, isGoodBoi bool) (*model.Dog, error) {
	id, err := gonanoid.New()
	if err != nil {
		panic("error creating id")
	}

	var createdDog = &model.Dog{ID: id, Name: name, IsGoodBoi: isGoodBoi}
	res := db.Create(&createdDog)

	return createdDog, res.Error
}

func GetDog(db *gorm.DB, id string) (*model.Dog, error) {
	var dog model.Dog
	res := db.First(&dog, "id = ?", id)

	return &dog, res.Error
}

func GetAllDogs(db *gorm.DB) ([]*model.Dog, error) {
	var dogs []*model.Dog

	res := db.Find(&dogs)

	return dogs, res.Error
}

func UpdateDogName(db *gorm.DB, id string, name string) (*model.Dog, error) {
	var dog model.Dog

	res := db.
		Model(&dog).
		// Model(&model.Dog{}). this works but returns empty data, w/out above Model called
		Clauses(clause.Returning{Columns: []clause.Column{}}).
		Where("id = ?", id).
		Update("name", name)

	return &dog, res.Error
}

func DeleteDog(db *gorm.DB, id string) (bool, error) {
	deleted := true

	res := db.Delete(&model.Dog{}, "id = ?", id)

	if res.Error != nil {
		deleted = false
		panic("Error deleting dog")
	}

	return deleted, res.Error
}
