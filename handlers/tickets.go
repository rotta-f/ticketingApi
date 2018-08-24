package handlers

import (
	"github.com/rotta-f/ticketingApi/database"
	"github.com/rotta-f/ticketingApi/datastructures"
	"github.com/rotta-f/ticketingApi/utils"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const (
	logHandlerTicket = "[HANDLER_TICKET] "
)

type ticketPayload struct {
	Title   string `json:"title"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

// Method: POST
// Route: /tickets/create
// Create a new ticket
func CreateTicket(w http.ResponseWriter, r *http.Request) {
	u := CtxGetValue(r, STORE_AUTH).(*datastructures.User)
	if u == nil {
		utils.WriteError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), "")
		return
	}

	payload := ticketPayload{}

	// Bind body with payload
	err := utils.BindJSON(r, &payload)
	if err != nil {
		log.Println(logHandlerTicket, "Fail to bind")
		utils.WriteError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err.Error())
		return
	}

	// Check payload
	if payload.Title == "" || payload.Message == "" {
		utils.WriteError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), "Title and message are required")
		return
	}

	// Create the ticket with first message
	t, err := database.CreateTicket(u.ID, payload.Title, payload.Message)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err.Error())
		return
	}
	utils.WriteJSON(w, t)
}

// Method: GET
// Route: /tickets
// Get all tickets without messages
func GetTickets(w http.ResponseWriter, r *http.Request) {
	// Get user id in Query params
	uID, err := strconv.ParseUint(r.URL.Query().Get("user"), 10, 64)
	if err != nil && r.URL.Query().Get("user") != "" {
		log.Println(logHandlerTicket, "ParseInt ", err)
		utils.WriteError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), "Invalid user ID")
		return
	}

	in := datastructures.Ticket{AuthorID: uint(uID)}

	ts, err := database.GetTickets(&in)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), "")
		return
	}
	utils.WriteJSON(w, ts)
}

// Method: GET
// Route: /tickets/{id}
// Get a single ticket with messages
func GetTicketById(w http.ResponseWriter, r *http.Request) {
	// Get ID in path
	urlT := strings.Split(r.URL.Path, "/")
	id, err := strconv.ParseUint(urlT[3], 10, 64)
	if err != nil {
		log.Println(logHandlerTicket, "ParseInt ", err)
		utils.WriteError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), "")
		return
	}

	t := &datastructures.Ticket{}
	t.ID = uint(id)
	t, err = database.GetTicket(t)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), "")
		return
	}
	utils.WriteJSON(w, t)
}

// Method: PATCH
// Route: /tickets/{id}
// Edit a ticket
func EditTicket(w http.ResponseWriter, r *http.Request) {
	u := CtxGetValue(r, STORE_AUTH).(*datastructures.User)
	if u == nil {
		utils.WriteError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), "")
		return
	}

	// Get ID in path
	urlT := strings.Split(r.URL.Path, "/")
	id, err := strconv.ParseUint(urlT[3], 10, 64)
	if err != nil {
		log.Println(logHandlerTicket, "ParseInt ", err)
		utils.WriteError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), "")
		return
	}

	payload := ticketPayload{}

	// Bind body with payload
	err = utils.BindJSON(r, &payload)
	if err != nil {
		log.Println(logHandlerTicket, "Fail to bind")
		utils.WriteError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err.Error())
		return
	}

	// Check payload
	if payload.Title == "" && payload.Message == "" {
		utils.WriteError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), "Title or status are required")
		return
	}


	// Check if user is support or if user is author of title
	t := &datastructures.Ticket{}
	if u.Type != datastructures.USER_TYPE_SUPPORT {
		t.ID = uint(id)
		t, err = database.GetTicket(t)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), "")
			return
		}
		if t.Author.ID != u.ID {
			utils.WriteError(w, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), "Client can only modify ticket he owns")
			return
		}
		payload.Status = ""
	}

	t = &datastructures.Ticket{}
	t.ID = uint(id)
	t.Status = payload.Status
	t.Title = payload.Title
	t, err = database.EditTicket(t)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), "")
		return
	}
	utils.WriteJSON(w, t)
}
