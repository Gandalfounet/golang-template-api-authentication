package userModule

import (
	"golang-email-api/modules/User/User/controllers/user"
	"github.com/gorilla/mux"
)

func Handlers(r *mux.Router) *mux.Router {
	
	// Auth route
	s := r.PathPrefix("/sendMail").Subrouter()
	s.HandleFunc("/", userModule.SendMail).Methods("POST")	
	return r
}