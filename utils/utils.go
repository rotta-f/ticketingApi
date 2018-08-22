package utils

import (
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const (
	logUtilsJson     = "[UTILS_JSON] "
	logUtilsToken    = "[UTILS_TOKEN] "
	logUtilsPassword = "[UTILS_TOKEN] "
)

func GeneratePassword() (string, error) {
	p := make([]byte, 4)
	i, err := rand.Read(p)
	if i != 4 || err != nil {
		log.Println(logUtilsPassword, err.Error())
		return "", err
	}
	return hex.EncodeToString(p), nil
}

func GenerateToken() (string, error) {
	token := make([]byte, 56)
	i, err := rand.Read(token)
	if i != 56 || err != nil {
		log.Println(logUtilsToken, err.Error())
		return "", err
	}

	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(time.Now().Unix()))

	b, err = bcrypt.GenerateFromPassword(append(b[:], token[:]...), bcrypt.DefaultCost)
	if i != 56 || err != nil {
		log.Println(logUtilsToken, err.Error())
		return "", err
	}

	return hex.EncodeToString(b), nil
}

func WriteError(w http.ResponseWriter, code int, err, msg string) {
	w.WriteHeader(code)
	WriteJSON(w, map[string]string{
		"error":   err,
		"message": msg,
	})
}

func BindJSON(r *http.Request, i interface{}) error {
	d := json.NewDecoder(r.Body)
	err := d.Decode(i)
	if err != nil {
		log.Println(logUtilsJson, "Decode JSON ", err)
		return err
	}
	return nil
}

func WriteJSON(w http.ResponseWriter, i interface{}) {
	j, err := json.Marshal(i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}
