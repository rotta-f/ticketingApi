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

type userCreatePayload struct {
	datastructures.User
	Password string `json:"password"`
}

const (
	logHandlerUser = "[HANDLER_USER] "
)

// This function create a user client/support depending on parameters 'create'
func userCreate(w http.ResponseWriter, r *http.Request, create func(user *datastructures.User) (*datastructures.User, error), t string) {
	// Check if user is a support
	u := CtxGetValue(r, STORE_AUTH).(*datastructures.User)
	if u == nil {
		utils.WriteError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), "")
		return
	}
	if u.Type != datastructures.USER_TYPE_SUPPORT {
		utils.WriteError(w, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), "Support required")
		return
	}

	var payload *userCreatePayload

	// Bind body with payload
	err := utils.BindJSON(r, &payload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err.Error())
		return
	}

	// Generate Password
	p, err := utils.GeneratePassword()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), "")
		return
	}

	// Init the new user
	payload.Firstname = t
	payload.Lastname = t
	payload.User.Password = p
	payload.Password = p
	u, err = create(&payload.User)
	payload.User = *u
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err.Error())
		return
	}
	utils.WriteJSON(w, payload)
}

// Method: POST
// Route: /users/create/support
// Create a support user
func UserCreateSupport(w http.ResponseWriter, r *http.Request) {
	userCreate(w, r, database.CreateUserSupport, "support")
}

// Method: POST
// Route: /users/create/client
// Create a support user
func UserCreateClient(w http.ResponseWriter, r *http.Request) {
	userCreate(w, r, database.CreateUserClient, "client")
}

// METHOD: PATCH
// Route: /users/edit/{id}
func UserUpdate(w http.ResponseWriter, r *http.Request) {
	u := CtxGetValue(r, STORE_AUTH).(*datastructures.User)
	if u == nil {
		utils.WriteError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), "")
		return
	}

	// Get ID in path
	urlT := strings.Split(r.URL.Path, "/")
	id, err := strconv.ParseUint(urlT[3], 10, 64)
	if err != nil {
		log.Println(logHandlerUser, "ParseInt ", err)
		utils.WriteError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), "")
		return
	}

	// Verify if client is right client
	if u.Type == datastructures.USER_TYPE_CLIENT && uint(id) != u.ID {
		utils.WriteError(w, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), "Can only modify yourself")
		return
	}

	var payload *userCreatePayload

	// Bind body with payload
	err = utils.BindJSON(r, &payload)
	if err != nil {
		log.Println(logHandlerUser, "Fail to bind")
		utils.WriteError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err.Error())
		return
	}

	if u.Type == datastructures.USER_TYPE_CLIENT {
		payload.Type = ""
	}
	payload.ID = uint(id)
	payload.User.Password = payload.Password
	err = database.UpdateUser(&payload.User)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Can't update user", err.Error())
		return
	}
	userUpdated, err := database.GetUserByID(payload.ID)
	if err != nil {
		log.Println(logHandlerUser, "GetUserByEmail ", err)
		utils.WriteError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), "")
		return
	}
	utils.WriteJSON(w, userUpdated)
}
