package database

import (
	"DB-DB/methods"
	"DB-DB/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// actually connect to the database!
	connection, err := gorm.Open(mysql.Open(methods.StringFromFile("dsn.txt")), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection

	DB.AutoMigrate(&models.DeliveryTicket{})     // automatically make sure that the tables in the DB are correct
	DB.AutoMigrate(&models.DeliveryTicketItem{}) // automatically make sure that the tables in the DB are correct
}
