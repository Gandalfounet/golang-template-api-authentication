package userController

import (
	"golang-template-api-authentication/modules/User/User/models"
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

func MagaAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Println("i am here")
}

func TestAPI(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API live and kicking"))
}



//CreateUser function -- create a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}
	json.NewDecoder(r.Body).Decode(user)

	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		err := ErrorResponse{
			Err: "Password Encryption  failed",
		}
		json.NewEncoder(w).Encode(err)
	}

	user.Password = string(pass)
	user.Role = "basic"

	rand.Seed(time.Now().UnixNano())

	validationToken := randSeq(25)

	contentMsg := utils.ContentLoginToken{Name: "Name", URL: "http://localhost/update/status/", Token: validationToken, Expiry: time.Now()}

	user.Status = "unverified"
	user.ResetToken = ""
	user.ResetTokenExpiracy = time.Now()
	user.ValidationToken = validationToken

	createdUser := db.Create(user)
	var errMessage = createdUser.Error

	if createdUser.Error != nil {
		fmt.Println(errMessage)
	}
	utils.Send(contentMsg)
	json.NewEncoder(w).Encode(createdUser)
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randSeq(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

func Me(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user")
	json.NewEncoder(w).Encode(user)
}

//FetchUser function
func FetchUsers(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user")
	fmt.Println(user)
	var users []models.User
	db.Preload("auths").Find(&users)

	json.NewEncoder(w).Encode(users)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	params := mux.Vars(r)
	var id = params["id"]
	db.First(&user, id)
	json.NewDecoder(r.Body).Decode(user)
	db.Save(&user)
	json.NewEncoder(w).Encode(&user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]
	var user models.User
	db.First(&user, id)
	db.Delete(&user)
	json.NewEncoder(w).Encode("User deleted")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]
	var user models.User
	db.First(&user, id)
	json.NewEncoder(w).Encode(&user)
}