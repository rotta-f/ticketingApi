package database

import (
	"github.com/rotta-f/ticketingApi/datastructures"
	"log"
)

const (
	logDatabaseMessage = "[DATABASE_MESSAGE] "
)

func AddMessageToTicket(t *datastructures.Ticket, userID uint, message string) (error) {
	m := &datastructures.Message{Text:message, AuthorID:userID}
	retAss := gDB.Model(t).Association("Messages").Append(m)
	if retAss.Error != nil {
		log.Println(logDatabaseMessage, "Add message ", retAss.Error)
		return retAss.Error
	}
	return nil
}