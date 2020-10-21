package linkycore

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitDB(o *LinkyCoreOptions) {
	// dbConfig := "root:project@(database)/project"
	db, err := gorm.Open("mysql", opts.DbURI+"?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect database")
	}

	if o.LogMode {
		db.LogMode(true)
	}

	DB = db
}

func CloseDB() {
	defer DB.Close()
}
