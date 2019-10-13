package routes

import (
	"golang-template-api-authentication/controllers"
	"golang-template-api-authentication/utils/auth"
	"net/http"

	"github.com/gorilla/mux"
)

func Handlers() *mux.Router {

	r := mux.NewRouter().StrictSlash(true)
	r.Use(CommonMiddleware)

	r.HandleFunc("/", controllers.TestAPI).Methods("GET")
	r.HandleFunc("/api", controllers.TestAPI).Methods("GET")
	r.HandleFunc("/register", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/reset/password", controllers.Login).Methods("POST")
	r.HandleFunc("/update/password/{token}", controllers.Login).Methods("PUT")
	r.HandleFunc("/update/status/{token}", controllers.Login).Methods("PUT")

	// Auth route
	s := r.PathPrefix("/me").Subrouter()
	s.Use(auth.JwtVerify)
	s.HandleFunc("/", controllers.Me).Methods("GET")
	s.HandleFunc("/update", controllers.Me).Methods("PUT")
	s.HandleFunc("/delete", controllers.Me).Methods("DELETE")
	
	// Admin route
	a := r.PathPrefix("/admin").Subrouter()
	a.Use(auth.JwtVerifyAdmin)
	a.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	a.HandleFunc("/users", controllers.FetchUsers).Methods("GET")
	a.HandleFunc("/user/{id}", controllers.GetUser).Methods("GET")
	a.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PUT")
	a.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")
	return r
}

// CommonMiddleware --Set content-type
func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		next.ServeHTTP(w, r)
	})
}