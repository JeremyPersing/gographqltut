package database

import (
	"os"

	"github.com/JeremyPersing/gographqltut/graph/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*
   // Create db logic, if using make sure yout connection string is configured to allow
   // Comment out if running multiple times - You may see an error otherwise
   db.Exec("CREATE DATABASE test_db")
   db.Exec("USE test_db")

*/

func connectToDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	dsn := os.Getenv("DB_CONNECTION_STRING")

	db, dbError := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if dbError != nil {
		panic("error connection to db")
	}
	return db
}

func performMigrations(db *gorm.DB) {
	db.AutoMigrate(&model.Dog{})
}

func InitializeDB() *gorm.DB {
	db := connectToDB()
	if db == nil {
		panic("db is nil")
	}
	performMigrations(db)
	return db
	// var createdDog = &model.Dog{ID: "1234523r245", Name: "Zip", IsGoodBoi: true}
	// db.Create(&createdDog)
}
