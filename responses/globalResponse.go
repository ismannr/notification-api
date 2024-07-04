package responses

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Status    int         `json:"status"`
	Message   string      `json:"message"`
	Timestamp string      `json:"timestamp"`
	Data      interface{} `json:"data,omitempty"`
}

func GlobalResponse(message string, status int, data interface{}, w http.ResponseWriter) {
	timestamp := time.Now().UTC().Format("2006-01-02T15:04:05.000Z07:00")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	res := Response{
		Status:    status,
		Message:   message,
		Timestamp: timestamp,
		Data:      data,
	}
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Println("error encoding json responses")
		return
	}
}
