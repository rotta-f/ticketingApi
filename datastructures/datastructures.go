package datastructures

import "github.com/jinzhu/gorm"

const (
	USER_TYPE_SUPPORT = "support"
	USER_TYPE_CLIENT  = "client"
)

type User struct {
	gorm.Model

	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email" gorm:";unique"`
	Type      string `json:"type"`
	Password  string `json:"-"`
}

const (
	TICKET_STATUS_OPENED        = "opened"
	TICKET_STATUS_PENDING_REPLY = "pending_reply"
	TICKET_STATUS_CLOSED        = "closed"
)

type Ticket struct {
	gorm.Model

	Title    string    `json:"title"`
	AuthorID uint      `json:"-"`
	Author   User      `json:"author"`
	Status   string    `json:"status"`
	Messages []Message `json:"messages,omitempty" gorm:"foreignkey:TicketID"`
}

type TicketArchive struct {
	Ticket
}

type Message struct {
	gorm.Model

	Text     string  `json:"text"`
	Author   User    `json:"author"`
	AuthorID uint    `json:"-"`
	Ticket   *Ticket `json:"ticket,omitempty"`
	TicketID uint    `json:"-"`
}

type Authentication struct {
	gorm.Model `json:"-"`

	Token  string `json:"token" gorm:";unique"`
	User   *User  `json:"user"`
	UserID int    `json:"-"`
}
