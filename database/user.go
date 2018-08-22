package database

import (
	"errors"
	"fmt"
	"github.com/rotta-f/ticketingApi/datastructures"
	"golang.org/x/crypto/bcrypt"
	"log"
)

const (
	logDatabaseUser = "[DATABASE_USER] "
)

func PrintDatabse() {
	fmt.Println("print")
}

func createUser(u *datastructures.User) (*datastructures.User, error) {
	if u.Firstname == "" {
		return nil, errors.New("Firstname empty")
	}
	if u.Lastname == "" {
		return nil, errors.New("Lastname empty")
	}
	if u.Email == "" {
		return nil, errors.New("Email empty")
	}
	if u.Type == "" {
		return nil, errors.New("Type empty")
	}
	if u.Password == "" {
		return nil, errors.New("Password empty")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Panic(logDatabaseUser, "Can't generate hash", err)
		return nil, err
	}

	u.Password = string(password)

	retDB := gDB.Create(u)
	if retDB.Error != nil {
		log.Println(logDatabaseUser, "Can't create user", retDB.Error)
		return nil, retDB.Error
	}

	return u, nil
}

func CreateUserClient(u *datastructures.User) (*datastructures.User, error) {
	u.Type = datastructures.USER_TYPE_CLIENT
	return createUser(u)
}

func CreateUserSupport(u *datastructures.User) (*datastructures.User, error) {
	u.Type = datastructures.USER_TYPE_SUPPORT
	return createUser(u)
}

func GetUserByEmail(email string) (*datastructures.User, error) {
	in := datastructures.User{Email: email}
	out := &datastructures.User{}
	if gDB.Where(&in).First(out).RecordNotFound() {
		err := errors.New("No user with email " + email)
		log.Println(logDatabaseUser, err)
		return nil, err
	}
	return out, nil
}
