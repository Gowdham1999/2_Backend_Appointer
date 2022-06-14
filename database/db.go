package database

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/entity"
)

var Connector *gorm.DB

func Connect(connectionString string) error {
	var err error

	Connector, err = gorm.Open("mysql", connectionString)

	if err != nil {
		return err
	}
	log.Println("Connection Successful")
	return nil
}

func Migrate(table *entity.Login) {
	Connector.AutoMigrate(&table)
	log.Println("Table Migration Successful")
}
