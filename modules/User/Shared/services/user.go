package userServices

import (
	"fmt"
	"time"
	"math/rand"
	"golang-template-api-authentication/modules/User/Shared/models"
	"golang.org/x/crypto/bcrypt"
	"golang-template-api-authentication/utils"
)

var db = utils.ConnectDB()

func CreateUser(user *models.User) (*models.User, error){
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(pass)
	user.Role = "basic"

	rand.Seed(time.Now().UnixNano())

	validationToken := randSeq(25)
	user.Status = "unverified"
	user.ResetToken = ""
	user.ResetTokenExpiry = time.Now()
	user.ValidationToken = validationToken

	userDb := db.Create(user)

	if userDb.Error != nil {
		fmt.Println(userDb.Error)
		return nil, userDb.Error
	}
	contentMsg := utils.ContentLoginToken{Name: "Name", URL: "http://localhost/update/status/", Token: validationToken, Expiry: time.Now()}
	go utils.Send(contentMsg, "resetPassword")
	return user, nil
}

func UpdateUser() {

}

func FindOne(user *models.User) (*models.User, error){
	fmt.Println(user.Email)
	fmt.Println(user.Password)
	userDb := &models.User{}

	if err := db.Where("Email = ?", user.Email).First(userDb).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	errf := bcrypt.CompareHashAndPassword([]byte(userDb.Password), []byte(user.Password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		fmt.Println(errf)
		return nil, errf
	}

	return userDb, nil
}

func GetUser() {

}

func DeleteUser() {

}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randSeq(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}