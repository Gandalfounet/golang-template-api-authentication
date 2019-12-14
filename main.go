package main

import (
	"fmt"
	"golang-email-api/routes"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	e := godotenv.Load(".env")

	if e != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println(e)

	port := os.Getenv("PORT")
	// Handle routes
	http.Handle("/", routes.Handlers())

	// serve
	log.Printf("Server up on port '%s'", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
