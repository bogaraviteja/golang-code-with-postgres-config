package database

import (
	"fmt"
	"post/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Db() *gorm.DB {

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=postgres password=ravi dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(&models.Person{})

	return db

}
