package userModule

import (
	"golang-email-api/utils"

	"encoding/json"
	"fmt"
	"net/http"

	//"github.com/gorilla/mux"
)

func Test(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(true)
}
func SendMail(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received mail")
	
	type User struct {
		Name     			string `json:"Name"`
		Email    			string `json:"Email"`
		Message   			string `json:"Message"`
		Date   			string `json:"Date"`
	}
	user := &User{}
	json.NewDecoder(r.Body).Decode(user)


	contentMsg := utils.ContentLoginToken{Name: user.Name, Email: user.Email, Message: user.Message, Date: user.Date}
	go utils.Send(contentMsg, "message")

	json.NewEncoder(w).Encode(contentMsg)
}
