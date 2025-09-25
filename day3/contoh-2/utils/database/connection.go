package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	// Read config from env
	DBHost := os.Getenv("DB_HOST")
	DBUser := os.Getenv("DB_USER")
	DBPassword := os.Getenv("DB_PASS")
	DBName := os.Getenv("DB_NAME")
	DBPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v TimeZone=Asia/Jakarta",
		DBHost,
		DBUser,
		DBPassword,
		DBName,
		DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error initiating database : ", err)
		os.Exit(1)
	}

	return db
}
