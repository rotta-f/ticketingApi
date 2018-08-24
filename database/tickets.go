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

func editTicket(tUpdate *datastructures.Ticket) error {
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

func ArchiveTicket(t *datastructures.Ticket) (error) {
	// Get the data to archive
	outT := datastructures.Ticket{}
	retDB := gDB.Where(t).Find(&outT)
	if retDB.Error != nil {
		log.Println(logDatabaseTicket, "ArchiveFind ", retDB.Error)
		return retDB.Error
	}
	outM := []datastructures.Message{}
	retDB = gDB.Model(t).Related(&outM)
	if retDB.Error != nil {
		log.Println(logDatabaseTicket, "ArchiveRelative ", retDB.Error)
		return retDB.Error
	}

	// Store the data in Archive
	retDB = gDBArchive.Create(outT)
	if retDB.Error != nil {
		log.Println(logDatabaseTicket, "ArchiveCreateTicket ", retDB.Error)
	}
	for _, elem := range outM {
		retDB = gDBArchive.Create(&elem)
		if retDB.Error != nil {
			log.Println(logDatabaseTicket, "ArchiveCreateMSG ", retDB.Error)
		}
	}

	// Delete old data
	retDB = gDB.Where("ID = ?", t.ID).Delete(datastructures.Ticket{})
	if retDB.Error != nil {
		log.Println(logDatabaseTicket, "ArchiveCreateTicket ", retDB.Error)
	}
	retDB = gDB.Where("Ticket_ID = ?", t.ID).Delete(datastructures.Message{})
	if retDB.Error != nil {
		log.Println(logDatabaseTicket, "ArchiveCreateTicket ", retDB.Error)
	}
	return nil
}
