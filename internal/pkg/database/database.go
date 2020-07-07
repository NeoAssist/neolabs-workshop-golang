package database

import (
	"fmt"
	"github.com/NeoAssist/neolabs-workshop-golang/internal/pkg/database/model"
	"github.com/NeoAssist/neolabs-workshop-golang/internal/utils"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // indirect
)

// New .
func New() *gorm.DB {
	dbHost := utils.GetEnv("DATABASE_HOST")
	dbPort := utils.GetEnv("DATABASE_PORT")
	dbUser := utils.GetEnv("DATABASE_USER")
	dbName := utils.GetEnv("DATABASE_NAME")
	dbPassword := utils.GetEnv("DATABASE_PASSWORD")

	connectionString := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPassword)

	db, err := gorm.Open("postgres", connectionString)

	if err != nil {
		fmt.Println("storage err: ", err)
	}

	db.DB().SetMaxIdleConns(3)
	db.LogMode(true)
	return db
}

// AutoMigrate .
func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.Poster{},
	)
}
