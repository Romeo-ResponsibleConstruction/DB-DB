package database

import (
	"DB-DB/models"
	"io"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// load setting from (so no passwords stored in plaintext in this file)
	// heavy use of https://stackoverflow.com/questions/1821811/how-to-read-write-from-to-a-file-using-go
	fi, err := os.Open("dsn.txt")
	if err != nil {
		panic(err)
	}
	// read to buffer
	buf := make([]byte, 1024)
	for {
		// read a chunk
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
	}
	// buffer -> string
	databaseSettings := string(buf)
	// close fi on exit and check for its returned error
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	// actually connect to the database!
	connection, err := gorm.Open(mysql.Open(databaseSettings), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection

	DB.AutoMigrate(&models.DeliveryTicket{})     // automatically make sure that the tables in the DB are correct
	DB.AutoMigrate(&models.DeliveryTicketItem{}) // automatically make sure that the tables in the DB are correct
}
