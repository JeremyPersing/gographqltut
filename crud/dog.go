package crud

import (
	"github.com/JeremyPersing/gographqltut/graph/model"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
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
