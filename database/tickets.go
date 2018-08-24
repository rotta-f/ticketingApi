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
	in := &datastructures.Ticket{}
	in.ID = t.ID
	return GetTicket(in)
}

func GetTicket(in *datastructures.Ticket) (*datastructures.Ticket, error) {
	out := &datastructures.Ticket{}
	retDB := gDB.Where(in).Preload("Author").Preload("Messages").Preload("Messages.Author").Find(out)
	if retDB.Error != nil {
		log.Println(logDatabaseTicket, "GetTicket ", retDB.Error)
		return nil, retDB.Error
	}
	return out, nil
}

func GetTickets(in *datastructures.Ticket) ([]datastructures.Ticket, error) {
	out := []datastructures.Ticket{}
	retDB := gDB.Where(in).Preload("Author").Find(&out)
	if retDB.Error != nil {
		return nil, retDB.Error
	}
	return out, nil
}

func editTicket(tUpdate *datastructures.Ticket) (error) {
	model := &datastructures.Ticket{}
	model.ID = tUpdate.ID

	retDB := gDB.Model(model).Update(tUpdate)
	if retDB.Error != nil {
		log.Println(logDatabaseTicket, "Update ", retDB.Error)
		return retDB.Error
	}
	return nil
}

func EditTicket(tUpdate *datastructures.Ticket) (*datastructures.Ticket, error) {
	err := editTicket(tUpdate)
	if err != nil {
		return nil, err
	}

	model := &datastructures.Ticket{}
	model.ID = tUpdate.ID
	return GetTicket(model)
}