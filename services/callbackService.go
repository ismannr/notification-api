package services

import (
	"api-notification/responses"
	"net/http"
)

func VerihubCallback(w http.ResponseWriter, r *http.Request) {
	sessionID := r.URL.Query().Get("session_id")
	status := r.URL.Query().Get("status")

	if sessionID == "" || status == "" {
		responses.GlobalResponse("Missing session_id or status parameter", http.StatusBadRequest, nil, w)
		return
	}

	// sangkutin ke db notif
}
