package authModule

import (
	"golang-template-api-authentication/modules/User/Shared/models"
	"golang-template-api-authentication/modules/User/Shared/services"
	"golang-template-api-authentication/utils"

	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"math/rand"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"

)

type ErrorResponse struct {
	Err string
}

type error interface {
	Error() string
}



var db = utils.ConnectDB()

//CreateUser function -- create a new user
func Register(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}
	json.NewDecoder(r.Body).Decode(user)

	userDb, err := userServices.CreateUser(user)
	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "An error occured", "code": 500}
		json.NewEncoder(w).Encode(resp)
		return
	}
	json.NewEncoder(w).Encode(userDb)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var token string
	user := &models.User{}

	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		//400 => Bad request
		var resp = map[string]interface{}{"status": false, "message": "Invalid request", "code": 400}
		json.NewEncoder(w).Encode(resp)
		return
	}

	userDb, err := userServices.FindOne(user)
	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Invalid credentials"}
		json.NewEncoder(w).Encode(resp)
		return
	}
	//u := resp["user"].(*models.User)
	if userDb.Status == "unverified" {
		//403 => Forbidden
		var resp = map[string]interface{}{"status": false, "message": "Not verified", "code": 403}
		json.NewEncoder(w).Encode(resp)
		return
	}

	token, err = userServices.GetToken(userDb)
	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "An error occured", "code": 500}
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
	type Email struct {
		Email string
	}

	email := &Email{}
	err := json.NewDecoder(r.Body).Decode(email)
	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
		json.NewEncoder(w).Encode(resp)
		return
	}

	user := &models.User{}

	if err := db.Where("Email = ?", email.Email).First(user).Error; err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Email address not found"}
		json.NewEncoder(w).Encode(resp)
		return
	}
	
	rand.Seed(time.Now().UnixNano())
	resetToken := randSeq(25)

	timein := time.Now().Add(time.Hour * 0 + time.Minute * 10 + time.Second * 0)

	user.ResetToken = resetToken
	user.ResetTokenExpiry = timein

	db.Save(&user)
	
	contentMsg := utils.ContentLoginToken{Name: "Name", URL: "http://localhost/update/password/", Token: resetToken, Expiry: timein}

	utils.Send(contentMsg, "resetPassword")
	response := true
	bolB, _ := json.Marshal(response)
	fmt.Println(string(bolB))

	json.NewEncoder(w).Encode(string(bolB))
}



func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	/*
	Password: "abc",
	Token: "abc"
	*/
	type test struct {
		Password string
	}
	passwordDatas := &test{}
	err := json.NewDecoder(r.Body).Decode(passwordDatas)

	if err != nil {
		fmt.Println(err)
		var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
		json.NewEncoder(w).Encode(resp)
		return
	}
	params := mux.Vars(r)
	token := params["token"]
	user := &models.User{}

	fmt.Println(token)
	if err := db.Where("reset_token = ?", token).First(user).Error; err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Invalid Token"}
		json.NewEncoder(w).Encode(resp)
		return
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(passwordDatas.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		err := ErrorResponse{
			Err: "Password Encryption  failed",
		}
		json.NewEncoder(w).Encode(err)
		return
	}

	user.Password = string(pass)

	db.Save(&user)

	contentMsg := utils.ContentLoginToken{Name: "Name", URL: "You changed your password", Token: "", Expiry: time.Now()}

	utils.Send(contentMsg, "resetPassword")
	response := true
	bolB, _ := json.Marshal(response)
	fmt.Println(string(bolB))

	json.NewEncoder(w).Encode(string(bolB))
}



func Validate(w http.ResponseWriter, r *http.Request) {
	/*
	Token : "abc"
	*/
	params := mux.Vars(r)
	token := params["token"]
	user := &models.User{}

	if err := db.Where("validation_token = ?", token).First(user).Error; err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Invalid Token"}
		json.NewEncoder(w).Encode(resp)
		return
	}

	user.Status = "verified"

	db.Save(&user)

	contentMsg := utils.ContentLoginToken{Name: "Name", URL: "Account confirmed", Token: "", Expiry: time.Now()}

	utils.Send(contentMsg, "resetPassword")

	
	type response1 struct {
	    Page   int
	    Fruits []string
	}
	type response2 struct {
	    Page   int      `json:"page"`
	    Fruits []string `json:"fruits"`
	}

	res1D := &response1{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res1B, _ := json.Marshal(res1D)
    fmt.Println(string(res1B))

    res2D := &response2{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res2B, _ := json.Marshal(res2D)

    fmt.Println(string(res2B))

    response := true
	bolB, _ := json.Marshal(response)
	fmt.Println(string(bolB))

	json.NewEncoder(w).Encode(string(res2B))
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randSeq(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}