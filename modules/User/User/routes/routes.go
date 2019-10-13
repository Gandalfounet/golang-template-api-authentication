package userRoutes

import (
	"golang-template-api-authentication/modules/User/User/controllers/user"
	"golang-template-api-authentication/modules/User/Authentication/utils/auth"

	"github.com/gorilla/mux"
)

func Handlers(r *mux.Router) *mux.Router {
	
	// Auth route
	s := r.PathPrefix("/me").Subrouter()
	s.Use(auth.JwtVerify)
	s.HandleFunc("/", userController.Me).Methods("GET")
	s.HandleFunc("/update", userController.Me).Methods("PUT")
	s.HandleFunc("/delete", userController.Me).Methods("DELETE")
	
	// Admin route
	a := r.PathPrefix("/admin").Subrouter()
	a.Use(auth.JwtVerifyAdmin)
	a.HandleFunc("/user", userController.CreateUser).Methods("POST")
	a.HandleFunc("/users", userController.FetchUsers).Methods("GET")
	a.HandleFunc("/user/{id}", userController.GetUser).Methods("GET")
	a.HandleFunc("/user/{id}", userController.UpdateUser).Methods("PUT")
	a.HandleFunc("/user/{id}", userController.DeleteUser).Methods("DELETE")
	
	return r
}