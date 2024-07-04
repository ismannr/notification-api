package auth

import (
	"api-notification/responses"
	"net/http"
	"os"
)

func ApiKeyAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("API-KEY")
		expectedApiKey := os.Getenv("API_KEY")

		if apiKey == "" {
			responses.GlobalResponse("API key is required", http.StatusUnauthorized, nil, w)
			return
		}

		if apiKey != expectedApiKey {
			responses.GlobalResponse("Invalid API key", http.StatusUnauthorized, nil, w)
			return
		}

		next.ServeHTTP(w, r)
	})
}
