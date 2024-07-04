package main

import (
	"api-notification/controllers"
	"api-notification/initializers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func init() {
	initializers.LoadEnv()
	initializers.DatabaseInit()
}
func main() {
	r := mux.NewRouter()
	controllers.NotificationController(r)
	controllers.CallbackController(r)
	controllers.SmsController(r)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Println(err.Error())
		return
	}
}
