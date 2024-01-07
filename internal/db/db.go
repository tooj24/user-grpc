package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func DatabaseSetup() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", "root", "root", "grpc_clean")
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatalf("There was error connecting to the database: %v", err)
	}

	return db
}
