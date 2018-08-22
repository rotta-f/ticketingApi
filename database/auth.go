package database

import (
	"github.com/rotta-f/ticketingApi/datastructures"
	"log"
)

const (
	logDatabaseAuth = "[DATABASE_AUTH] "
)

func AddAuthToken(a *datastructures.Authentication) error {
	a.UserID = int(a.User.ID)
	retDB := gDB.Create(a)
	if retDB.Error != nil {
		log.Println(logDatabaseAuth, "Can't add token", retDB.Error)
		return retDB.Error
	}
	return nil
}

func GetAuthToken(t string) *datastructures.Authentication {
	in := &datastructures.Authentication{Token: t}
	out := &datastructures.Authentication{}
	if gDB.Where(in).Preload("User").First(out).RecordNotFound() {
		return nil
	}
	return out
}
