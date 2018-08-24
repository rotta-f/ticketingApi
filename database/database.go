package database

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/rotta-f/ticketingApi/datastructures"
)

const (
	logDatabaseInit = "[DATABASE_INIT] "
)

var gDB *gorm.DB
var gDBArchive *gorm.DB

func init() {
	db, err := gorm.Open("sqlite3", "ticketing.db")
	if err != nil {
		log.Fatal(logDatabaseInit, err)
	}

	db.SingularTable(true)
	db.AutoMigrate(&datastructures.Authentication{}, &datastructures.User{}, &datastructures.Message{}, &datastructures.Ticket{})
	gDB = db

	dbA, err := gorm.Open("sqlite3", "ticketingArchive.db")
	if err != nil {
		log.Fatal(logDatabaseInit, err)
	}
	dbA.SingularTable(true)
	dbA.AutoMigrate(&datastructures.Message{}, &datastructures.Ticket{})
	gDBArchive = dbA

	// Create admin user if new database
	var c int
	retDB := db.Model(&datastructures.User{}).Count(&c)
	if retDB.Error != nil {
		log.Fatal(logDatabaseInit, retDB.Error)
	}
	if c == 0 {
		_, err := CreateUserSupport(&datastructures.User{Firstname: "admin", Lastname: "admin", Email: "admin@ticket.lu", Password: "admin"})
		if err != nil {
			log.Fatal(logDatabaseInit, err)
		}
	}
}
