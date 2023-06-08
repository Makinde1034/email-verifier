package routes

import (
	"github.com/gorilla/mux"
	"github.com/Makinde1034/email-verifier/controllers"
)


func RegisterRoutes() *mux.Router{
	router := mux.NewRouter()
	router.HandleFunc("/verify-email",controllers.VerifyEmail).Methods("POST")
	return router
}