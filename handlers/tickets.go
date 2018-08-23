package handlers

import (
	"net/http"
	"github.com/rotta-f/ticketingApi/datastructures"
	"github.com/rotta-f/ticketingApi/utils"
	"log"
	"github.com/rotta-f/ticketingApi/database"
)

type ticketPayload struct {
	Title   string `json:"title"`
	Message string `json:"message"`
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
		log.Println(logHandlerUser, "Fail to bind")
		utils.WriteError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err.Error())
		return
	}

	if payload.Title == "" || payload.Message == "" {
		utils.WriteError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), "Title and message are required")
		return
	}

	t, err := database.CreateTicket(u.ID, payload.Title, payload.Message)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err.Error())
		return
	}
	utils.WriteJSON(w, t)
}