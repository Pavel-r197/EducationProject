package postgres

import (
	"EducationProject/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func OpenDb(dsn string) *gorm.DB {

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func AutoMigrate(db *gorm.DB) {
	if err := db.AutoMigrate(&domain.Task{}, &domain.User{}); err != nil {
		log.Fatal(err)
	}

}
