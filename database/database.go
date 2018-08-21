package database

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/rotta-f/ticketingApi/datastructures"
)

const (
	LOG_DATABASE_INIT = "[DATABASE_INIT] "
)

var gDB *gorm.DB

func init() {
	db, err := gorm.Open("sqlite3", "ticketing.db")
	defer db.Close()

	if err != nil {
		log.Fatal(LOG_DATABASE_INIT, err)
	}

	db.SingularTable(true)
	db.AutoMigrate(&datastructures.User{}, &datastructures.Ticket{}, &datastructures.TicketArchive{}, &datastructures.Message{})

	gDB = db

	// Create admin user if new database
	var c int
	retDB := db.Model(&datastructures.User{}).Count(&c)
	if retDB.Error != nil {
		log.Fatal(LOG_DATABASE_INIT, retDB.Error)
	}
	if c == 0 {
		_, err := CreateUserSupport(&datastructures.User{Firstname: "admin", Lastname: "admin", Email: "admin@ticket.lu", Password: "admin"})
		if err != nil {
			log.Fatal(LOG_DATABASE_INIT, err)
		}
	}
}
