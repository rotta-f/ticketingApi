package database

import (
	"github.com/rotta-f/ticketingApi/datastructures"
	"log"
)

const (
	logDatabaseTicket = "[DATABASE_TICKET] "
)

func CreateTicket(userID uint, title, message string) (*datastructures.Ticket, error) {
	t := &datastructures.Ticket{Title: title, AuthorID: userID, Status: datastructures.TICKET_STATUS_OPENED}
	retDB := gDB.Create(t)
	if retDB.Error != nil {
		log.Println(logDatabaseTicket, "Create ticket ", retDB.Error)
		return nil, retDB.Error
	}

	err := AddMessageToTicket(t, userID, message)
	if err != nil {
		return nil, retDB.Error
	}
	return GetTicket(t)
}

func GetTicket(in *datastructures.Ticket) (*datastructures.Ticket, error) {
	out := &datastructures.Ticket{}
	retDB := gDB.Model(in).Preload("Author").Preload("Messages").Preload("Messages.Author").Find(out)
	if retDB.Error != nil {
		return nil, retDB.Error
	}
	return out, nil
}