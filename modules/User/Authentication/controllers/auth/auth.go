package authModule

import (
	"golang-template-api-authentication/modules/User/Shared/models"
	"golang-template-api-authentication/modules/User/Shared/services"
	"golang-template-api-authentication/utils"

	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"

)

var db = utils.ConnectDB()

//CreateUser function -- create a new user
func Register(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}
	json.NewDecoder(r.Body).Decode(user)

	userDb, err := userServices.CreateUser(user)
	if err != nil {
		resp, err := utils.GetError(500, "fr")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(resp)
		json.NewEncoder(w).Encode(resp)
		return
	}

	contentMsg := utils.ContentLoginToken{Name: "Name", URL: "http://localhost/update/status/", Token: userDb.ValidationToken}
	go utils.Send(contentMsg, "loginToken")

	json.NewEncoder(w).Encode(userDb)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var token string
	user := &models.User{}

	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		resp, err := utils.GetError(400, "en")
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	userDb, err := userServices.FindByEmail(user)
	if err != nil {
		resp, err := utils.GetError(404, "en")
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	errPw := userServices.CheckPassword(userDb, user.Password)
	if errPw != nil {
		resp, err := utils.GetError(404, "en")
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	//u := resp["user"].(*models.User)
	if userDb.Status == "unverified" {
		resp, err := utils.GetError(403, "en")
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	token, err = userServices.GetToken(userDb)
	if err != nil {
		resp, err := utils.GetError(500, "en")
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	var resp = map[string]interface{}{"status": false, "message": "logged in"}
	resp["token"] = token //Store the token in the response
	resp["user"] = userDb

	json.NewEncoder(w).Encode(resp)
}

func Me(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user")
	json.NewEncoder(w).Encode(user)
}



func ResetPassword(w http.ResponseWriter, r *http.Request) {
	/*
	Email: "abc"
	*/

	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		resp, err := utils.GetError(400, "en")
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	userDb, err := userServices.FindByEmail(user)
	if err != nil {
		resp, err := utils.GetError(404, "en")
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}
		json.NewEncoder(w).Encode(resp)
		return
	}
	
	resetToken, timein := userServices.GetTokensPassword()
	userDb.ResetToken = resetToken
	userDb.ResetTokenExpiry = timein
	userServices.UpdateUser(userDb)
	
	contentMsg := utils.ContentLoginToken{Name: "Name", URL: "http://localhost/update/password/", Token: resetToken, Expiry: timein}

	go utils.Send(contentMsg, "resetPassword")
	var resp = map[string]interface{}{"status": 200, "message": "Success"}
	json.NewEncoder(w).Encode(resp)
}


func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	/*
	Password: "abc", as POST
	token: "abc" as get param
	*/
	type test struct {
		Password string
	}
	passwordDatas := &test{}
	err := json.NewDecoder(r.Body).Decode(passwordDatas)
	if err != nil {
		resp, err := utils.GetError(400, "en")
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	params := mux.Vars(r)
	token := params["token"]
	userDb, err := userServices.FindByToken(token)
	if err != nil {
		resp, err := utils.GetError(404, "en")
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	//Check token expiry
	diff := time.Now().Sub(userDb.ResetTokenExpiry)
	if(diff.Seconds() > 0) {
		resp, err := utils.GetError(401, "en")
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(passwordDatas.Password), bcrypt.DefaultCost)
	if err != nil {
		resp, err := utils.GetError(500, "en")
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	userDb.Password = string(pass)

	userServices.UpdateUser(userDb)

	contentMsg := utils.ContentLoginToken{Name: "Name"}

	go utils.Send(contentMsg, "updatePassword")
	var resp = map[string]interface{}{"status": 200, "message": "Success"}
	json.NewEncoder(w).Encode(resp)
}



func Validate(w http.ResponseWriter, r *http.Request) {
	/*
	Token : "abc"
	*/
	params := mux.Vars(r)
	token := params["token"]

	userDb, err := userServices.FindByValidationToken(token)
	if err != nil {
		resp, err := utils.GetError(404, "en")
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	userDb.Status = "verified"

	userServices.UpdateUser(userDb)

	contentMsg := utils.ContentLoginToken{Name: "Name"}

	go utils.Send(contentMsg, "validatePassword")

	var resp = map[string]interface{}{"status": 200, "message": "Success"}
	json.NewEncoder(w).Encode(resp)
}