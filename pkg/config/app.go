package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() {
	// use the mysql driver Dialector required by GORM v2
	dsn := "root:Mahi9313711905@@tcp(localhost:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
