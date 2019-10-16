package userModule

import (
	"golang-template-api-authentication/modules/User/User/controllers/user"
	"golang-template-api-authentication/modules/User/Shared/utils/auth"

	"github.com/gorilla/mux"
)

func Handlers(r *mux.Router) *mux.Router {
	
	// Auth route
	s := r.PathPrefix("/me").Subrouter()
	s.Use(auth.JwtVerify)
	s.HandleFunc("/", userModule.Me).Methods("GET")
	s.HandleFunc("/update", userModule.Me).Methods("PUT")
	s.HandleFunc("/delete", userModule.Me).Methods("DELETE")
	
	// Admin route
	a := r.PathPrefix("/admin").Subrouter()
	a.Use(auth.JwtVerifyAdmin)
	a.HandleFunc("/user", userModule.FetchUsers).Methods("POST")
	a.HandleFunc("/users", userModule.FetchUsers).Methods("GET")
	a.HandleFunc("/user/{id}", userModule.GetUser).Methods("GET")
	a.HandleFunc("/user/{id}", userModule.UpdateUser).Methods("PUT")
	a.HandleFunc("/user/{id}", userModule.DeleteUser).Methods("DELETE")
	
	return r
}