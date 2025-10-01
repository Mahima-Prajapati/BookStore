package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() {
	// load env variables
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	// use the mysql driver Dialector required by GORM v2'
	dbAdd := os.Getenv("DB_ADDR")
	dbPass := os.Getenv("DB_PASS")

	dsn := fmt.Sprintf("root:%s@tcp(%s)/bookstore?charset=utf8mb4&parseTime=True&loc=Local", dbPass, dbAdd)
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
