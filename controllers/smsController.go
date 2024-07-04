package controllers

import (
	"api-notification/auth"
	"api-notification/services"
	"github.com/gorilla/mux"
)

func SmsController(r *mux.Router) {
	api := r.PathPrefix("/sms").Subrouter()
	api.Use(auth.ApiKeyAuth)
	api.HandleFunc("/send-otp", services.SendOtpBySMS).Methods("POST")
	api.HandleFunc("/verify", services.VerifySmsOtp).Methods("POST")
}
