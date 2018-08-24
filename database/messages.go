package database

import (
	"github.com/rotta-f/ticketingApi/datastructures"
	"log"
)

const (
	logDatabaseMessage = "[DATABASE_MESSAGE] "
)

func AddMessageToTicket(t *datastructures.Ticket, userID uint, message string) error {
	m := &datastructures.Message{Text: message, AuthorID: userID}
	retAss := gDB.Model(t).Association("Messages").Append(m)
	if retAss.Error != nil {
		log.Println(logDatabaseMessage, "Add message ", retAss.Error)
		return retAss.Error
	}
	return nil
}

func NewMessageToTicket(t *datastructures.Ticket, userID uint, message string) (*datastructures.Message, error) {
	m := &datastructures.Message{Text: message, AuthorID: userID}
	retAss := gDB.Model(t).Association("Messages").Append(m)
	if retAss.Error != nil {
		log.Println(logDatabaseMessage, "New message ", retAss.Error)
		return nil, retAss.Error
	}
	t.Status = datastructures.TICKET_STATUS_PENDING_REPLY
	err := editTicket(t)
	if err != nil {
		log.Println(logDatabaseMessage, "Update pending reply", err)
		return nil, err
	}
	return m, nil
}

func GetMessage(in *datastructures.Message) (*datastructures.Message, error) {
	out := &datastructures.Message{}
	retDB := gDB.Where(in).Preload("Author").Preload("Ticket").Preload("Ticket.Author").Find(out)
	if retDB.Error != nil {
		log.Println(logDatabaseMessage, "GetMessage ", retDB.Error)
		return nil, retDB.Error
	}
	return out, nil
}

func GetTicketMessages(in *datastructures.Ticket) ([]datastructures.Message, error) {
	out := []datastructures.Message{}
	retDB := gDB.Model(in).Preload("Author").Related(&out)
	if retDB.Error != nil {
		log.Println(logDatabaseMessage, "GetTicketMessages ", retDB.Error)
		return nil, retDB.Error
	}
	return out, nil
}

func EditMessage(in *datastructures.Message) (*datastructures.Message, error) {
	model := datastructures.Message{}
	model.ID = in.ID
	retDB := gDB.Model(&model).Update(in)
	if retDB.Error != nil {
		log.Println(logDatabaseMessage, "GetTicketMessages ", retDB.Error)
		return nil, retDB.Error
	}
	return GetMessage(&model)
}