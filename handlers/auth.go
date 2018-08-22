package handlers

import (
	"github.com/rotta-f/ticketingApi/database"
	"github.com/rotta-f/ticketingApi/datastructures"
	"github.com/rotta-f/ticketingApi/utils"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

const (
	STORE_AUTH = "auth"
)

type authRequestPayload struct {
	*datastructures.User
	Password string `json:"password"`
}

type authResponsePayload struct {
	Token string               `json:"token"`
	User  *datastructures.User `json:"user"`
}

// Create access token and store the token in memory linked to the uer
func initConnection(u *datastructures.User) (string, error) {
	t, err := utils.GenerateToken()
	if err != nil {
		return "", err
	}

	a := &datastructures.Authentication{Token: t, User: u}
	database.AddAuthToken(a)

	return t, nil
}

// Method: POST
// Route: /auth/login
// Authentication via email / password
func AuthLogin(w http.ResponseWriter, r *http.Request) {
	var payload authRequestPayload

	err := utils.BindJSON(r, &payload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err.Error())
		return
	}

	if payload.Password == "" || payload.Email == "" {
		utils.WriteError(w, http.StatusBadRequest, "Missing fields", "Email or Password are empty")
		return
	}

	u, err := database.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, "Bad credential", "Email or Password wrong")
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(payload.Password)); err != nil {
		utils.WriteError(w, http.StatusUnauthorized, "Bad credential", "Email or Password wrong")
		return
	}

	t, err := initConnection(u)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), "")
		return
	}

	utils.WriteJSON(w, authResponsePayload{Token: t, User: u})
}

// Method: POST
// Route: /auth/signup
// Create a new client user
func AuthSignup(w http.ResponseWriter, r *http.Request) {
	var payload authRequestPayload

	err := utils.BindJSON(r, &payload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err.Error())
		return
	}

	payload.User.Password = payload.Password
	payload.User, err = database.CreateUserClient(payload.User)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err.Error())
		return
	}

	t, err := initConnection(payload.User)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), "")
		return
	}

	utils.WriteJSON(w, authResponsePayload{Token: t, User: payload.User})
}
