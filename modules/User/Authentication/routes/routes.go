package authRoutes

import (
	"golang-template-api-authentication/modules/User/User/controllers/user"
	"golang-template-api-authentication/modules/User/Authentication/controllers/auth"

	"github.com/gorilla/mux"
)

func Handlers(r *mux.Router) *mux.Router {

	r.HandleFunc("/", userController.TestAPI).Methods("GET")
	r.HandleFunc("/api", userController.TestAPI).Methods("GET")
	r.HandleFunc("/register", userController.CreateUser).Methods("POST")
	r.HandleFunc("/login", authController.Login).Methods("POST")
	r.HandleFunc("/reset/password", authController.ResetPassword).Methods("POST")
	r.HandleFunc("/update/password/{token}", authController.UpdatePassword).Methods("PUT")
	r.HandleFunc("/update/status/{token}", authController.Validate).Methods("PUT")
	
	return r
}