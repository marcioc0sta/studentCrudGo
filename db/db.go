package db

import (
	"log"

	"github.com/marcioc0sta/gin-api-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectWithDatabase() {
	connectionStr := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectionStr))
	if err != nil {
			log.Panic("Error on connecting to database")
	}
	DB.AutoMigrate(&models.Student{})
}