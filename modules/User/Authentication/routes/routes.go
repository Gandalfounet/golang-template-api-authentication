package authModule

import (
	"golang-template-api-authentication/modules/User/User/controllers/user"
	"golang-template-api-authentication/modules/User/Authentication/controllers/auth"

	"github.com/gorilla/mux"
)

func Handlers(r *mux.Router) *mux.Router {

	r.HandleFunc("/", userModule.TestAPI).Methods("GET")
	r.HandleFunc("/api", userModule.TestAPI).Methods("GET")
	r.HandleFunc("/register", userModule.CreateUser).Methods("POST")
	r.HandleFunc("/login", authModule.Login).Methods("POST")
	r.HandleFunc("/reset/password", authModule.ResetPassword).Methods("POST")
	r.HandleFunc("/update/password/{token}", authModule.UpdatePassword).Methods("PUT")
	r.HandleFunc("/validate/{token}", authModule.Validate).Methods("GET")
	
	return r
}