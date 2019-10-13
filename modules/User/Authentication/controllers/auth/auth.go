package authController

import (
	"golang-template-api-authentication/modules/User/User/models"
	"golang-template-api-authentication/utils"

	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"math/rand"

	jwt "github.com/dgrijalva/jwt-go"
	//"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"

	"os"
	"github.com/joho/godotenv"
)

type ErrorResponse struct {
	Err string
}

type error interface {
	Error() string
}

var db = utils.ConnectDB()

func Login(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := FindOne(user.Email, user.Password)
	json.NewEncoder(w).Encode(resp)
}

func FindOne(email, password string) map[string]interface{} {
	user := &models.User{}

	if err := db.Where("Email = ?", email).First(user).Error; err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Email address not found"}
		return resp
	}
	expiresAt := time.Now().Add(time.Minute * 100000).Unix()

	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		var resp = map[string]interface{}{"status": false, "message": "Invalid login credentials. Please try again"}
		return resp
	}

	tk := &models.Token{
		UserID: user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Role: user.Role,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	//Get the secret key from .env
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		fmt.Println("Error loading .env file")
	}
	secretJwt := os.Getenv("secretJwt")

	tokenString, error := token.SignedString([]byte(secretJwt))
	if error != nil {
		fmt.Println(error)
	}

	var resp = map[string]interface{}{"status": false, "message": "logged in"}
	resp["token"] = tokenString //Store the token in the response
	resp["user"] = user
	return resp
}

func Me(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user")
	json.NewEncoder(w).Encode(user)
}

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	/*
	Email: "abc"
	*/
	email := &models.Email{}
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
	resetTokenExpiracy := ""

	user.ResetToken = resetToken
	user.ResetTokenExpiracy = resetTokenExpiracy

	db.Save(&user)
	utils.Send(resetToken)
	json.NewEncoder(w).Encode(&user)
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	/*
	Password: "abc",
	Token: "abc"
	*/
}

func Validate(w http.ResponseWriter, r *http.Request) {
	/*
	Token : "abc"
	*/
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randSeq(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}