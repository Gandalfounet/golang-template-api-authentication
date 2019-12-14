package userModule

import (
	"golang-email-api/modules/User/User/controllers/user"
	"github.com/gorilla/mux"
)

func Handlers(r *mux.Router) *mux.Router {
	
	// Auth route
	r.HandleFunc("/", userModule.Test).Methods("GET")	
	r.HandleFunc("/sendMail", userModule.SendMail).Methods("POST")	
	return r
}