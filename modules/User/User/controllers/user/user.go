package userModule

import (
	"golang-email-api/utils"

	//"encoding/json"
	//"fmt"
	"net/http"

	//"github.com/gorilla/mux"
)


func SendMail(w http.ResponseWriter, r *http.Request) {

	contentMsg := utils.ContentLoginToken{Name: "Name", URL: "http://localhost:8080/auth/validate/"}
	go utils.Send(contentMsg, "loginToken")
}
