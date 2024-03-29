package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDb() *gorm.DB {

	db,err := gorm.Open(postgres.Open(os.Getenv("DBS")), &gorm.Config{})

	if err != nil {
		fmt.Println("database not loaded:", err)
	}

	fmt.Println("Succesfully Connected db")

	return db
}