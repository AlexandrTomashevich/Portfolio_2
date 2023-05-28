package friend_service

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Error string `json:"Error"`
}

// для отправки ошибки клиенту
func writeError(w http.ResponseWriter, msg string, status int) {
	w.WriteHeader(status)
	e := &Error{Error: msg}
	byteMessage, _ := json.Marshal(e)
	_, err := w.Write(byteMessage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
