package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"time"
	"encoding/binary"
	"golang.org/x/crypto/bcrypt"
	"encoding/hex"
)

const (
	logUtilsJson = "[UTILS_JSON] "
	logUtilsToken = "[UTILS_TOKEN] "
)

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
