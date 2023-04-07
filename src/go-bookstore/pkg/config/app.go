package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	data, err := gorm.Open("mysql", "mysql-db:password@/simpleCrudGo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db = data
}

func GetDb() *gorm.DB {
	return db
}
