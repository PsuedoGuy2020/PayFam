package Db

import (
	"PayFam/internal/models/dao"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func InitDatabase(host, user, password, dbname string, port string) (*gorm.DB, error) {
	dbURL := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&dao.Video{})
	if err != nil {
		log.Fatalf("Failed to migrate the database: %v", err)
	}

	return db, nil
}
