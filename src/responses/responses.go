package responses

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)

	if error := json.NewEncoder(w).Encode(data); error != nil {
		log.Fatal(error)
	}
}

func Error(w http.ResponseWriter, statusCode int, error error) {
	JSON(w, statusCode, struct {
		Erro      string    `json:"erro"`
		Timestamp time.Time `json:"timestamp"`
	}{
		Erro:      error.Error(),
		Timestamp: time.Now(),
	})
}
