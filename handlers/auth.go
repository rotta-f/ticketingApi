package handlers

import (
	"github.com/rotta-f/ticketingApi/database"
	"github.com/rotta-f/ticketingApi/datastructures"
	"github.com/rotta-f/ticketingApi/utils"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"github.com/rotta-f/ticketingApi/router"
)

type authLoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type authPayload struct {
	Token string              `json:"token"`
	User  *datastructures.User `json:"user"`
}

func AuthLogin(w http.ResponseWriter, r *http.Request) {
	var payload authLoginPayload

	err := utils.BindJSON(r, &payload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Bad payload", err.Error())
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

	t, err := utils.GenerateToken()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), "")
		return
	}

	router.InitStore(t)
	router.GetStore(t).Set("user", u)
	var uu *datastructures.User
	uu = router.GetStore(t).Get("user").(*datastructures.User)
	println(uu.Email)


	utils.WriteJSON(w, authPayload{Token: t, User: u})
}
