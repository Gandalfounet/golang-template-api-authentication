package authModule

import (
	"golang-template-api-authentication/modules/User/User/controllers/user"
	"golang-template-api-authentication/modules/User/Authentication/controllers/auth"

	"github.com/gorilla/mux"
)

func Handlers(r *mux.Router) *mux.Router {

	r.HandleFunc("/", userModule.TestAPI).Methods("GET")
	r.HandleFunc("/api", userModule.TestAPI).Methods("GET")
	

	// Auth route
	p := r.PathPrefix("/auth").Subrouter()
	p.HandleFunc("/register", authModule.Register).Methods("POST")
	p.HandleFunc("/login", authModule.Login).Methods("POST")
	p.HandleFunc("/reset/password", authModule.ResetPassword).Methods("POST")
	p.HandleFunc("/update/password/{token}", authModule.UpdatePassword).Methods("PUT")
	p.HandleFunc("/validate/{token}", authModule.Validate).Methods("GET")

	return r
}