package handlers

import (
	"net/http"
	"strings"
	"strconv"
	"github.com/rotta-f/ticketingApi/datastructures"
	"log"
	"github.com/rotta-f/ticketingApi/utils"
	"github.com/rotta-f/ticketingApi/database"
)

const (
	logHandlerMessage = "[HANDLER_MESSAGE] "
)

// Method: POST
// Route: /messages/ticket/{id_ticket}
// Add a message to a ticket
func NewMessageToTicket(w http.ResponseWriter, r *http.Request) {
	u := CtxGetValue(r, STORE_AUTH).(*datastructures.User)
	if u == nil {
		utils.WriteError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), "")
		return
	}

	// Get ID in path
	urlT := strings.Split(r.URL.Path, "/")
	id, err := strconv.ParseUint(urlT[4], 10, 64)
	if err != nil {
		log.Println(logHandlerMessage, "ParseInt ", err)
		utils.WriteError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), "")
		return
	}

	m := &datastructures.Message{}
	// Bind body with payload
	err = utils.BindJSON(r, m)
	if err != nil {
		log.Println(logHandlerMessage, "Fail to bind")
		utils.WriteError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err.Error())
		return
	}

	if m.Text == "" {
		utils.WriteError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), "Text required")
		return
	}

	// Check if ticket is not closed before update
	t := &datastructures.Ticket{}
	t.ID = uint(id)
	t, err = database.GetTicket(t)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err.Error())
		return
	}
	if t.Status == datastructures.TICKET_STATUS_CLOSED {
		utils.WriteError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), "Cannot add message on closed tickets")
		return
	}

	// Insert new message to ticket
	m, err = database.NewMessageToTicket(t, u.ID, m.Text)
	if err != nil {
		log.Println(logHandlerMessage, err)
		utils.WriteError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), "")
		return
	}

	// Get the message for response
	m = &datastructures.Message{}
	m.ID = m.ID
	m, err = database.GetMessage(m)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), "")
		return
	}

	utils.WriteJSON(w, m)
}

func GetTicketMessages(w http.ResponseWriter, r *http.Request) {
	// Get ID in path
	urlT := strings.Split(r.URL.Path, "/")
	id, err := strconv.ParseUint(urlT[4], 10, 64)
	if err != nil {
		log.Println(logHandlerMessage, "ParseInt ", err)
		utils.WriteError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), "")
		return
	}

	t := &datastructures.Ticket{}
	t.ID = uint(id)
	ms, err := database.GetTicketMessages(t)
	if err != nil {
		log.Println(logHandlerMessage, err)
		utils.WriteError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), "")
		return
	}

	utils.WriteJSON(w, ms)
}