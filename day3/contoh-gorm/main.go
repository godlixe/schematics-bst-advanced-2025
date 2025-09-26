package main

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	DBHost := "localhost"
	DBUser := "postgres"
	DBPassword := "root"
	DBName := "blogs"
	DBPort := "5433"

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

type User struct {
	gorm.Model
	Username string
	Age      int
}

func main() {
	db := ConnectDatabase()

	db.AutoMigrate(&User{})

	elsi := User{
		Username: "elsi",
		Age:      23,
	}

	// menambahkan data
	tx := db.Create(&elsi)
	if tx.Error != nil {
		fmt.Println(tx.Error)
	}

	// mengupdate/mengubah data
	tx = db.Debug().Where("id = ?", 1).Updates(&elsi)
	if tx.Error != nil {
		fmt.Println(tx.Error)
	}

	var user []User

	// mencari data
	tx = db.Debug().Where("id = 1").Find(&user)
	if tx.Error != nil {
		fmt.Println(tx.Error)
	}

	// menghapus data
	tx = db.Debug().Where("id = 1").Delete(&User{})
	if tx.Error != nil {
		fmt.Println(tx.Error)
	}

	tx = db.Debug().Where("id = 1").Find(&user)
	if tx.Error != nil {
		fmt.Println(tx.Error)
	}

	fmt.Println(user)
}
