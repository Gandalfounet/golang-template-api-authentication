package userServices

import (
	"fmt"
	"time"
	"math/rand"
	"golang-template-api-authentication/modules/User/Shared/models"
	"golang.org/x/crypto/bcrypt"
	"golang-template-api-authentication/utils"
)

func CreateUser(user *models.User) (*models.User, error){
	db := utils.GetDB()
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
	
	return user, nil
}

func UpdateUser(user *models.User) {
	db := utils.GetDB()
	db.Save(&user)
}

func GetTokensPassword() (string, time.Time) {
	resetToken := GenerateToken()
	timein := time.Now().Add(time.Hour * 0 + time.Minute * 10 + time.Second * 0)
	
	return resetToken, timein
}
 
func GenerateToken() string {
	rand.Seed(time.Now().UnixNano())
	resetToken := randSeq(25)
	return resetToken
}

func FindByEmail(user *models.User) (*models.User, error) {
	db := utils.GetDB()
	userDb := &models.User{}

	if err := db.Where("Email = ?", user.Email).First(userDb).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return userDb, nil
}

func FindByToken(token string) (*models.User, error) {
	db := utils.GetDB()
	userDb := &models.User{}

	if err := db.Where("reset_token = ?", token).First(userDb).Error; err != nil {
		// var resp = map[string]interface{}{"status": false, "message": "Invalid Token"}
		// json.NewEncoder(w).Encode(resp)
		// return
		return nil, err
	}

	return userDb, nil
}

func FindByValidationToken(token string) (*models.User, error) {
	db := utils.GetDB()
	userDb := &models.User{}

	if err := db.Where("validation_token = ?", token).First(userDb).Error; err != nil {
		// var resp = map[string]interface{}{"status": false, "message": "Invalid Token"}
		// json.NewEncoder(w).Encode(resp)
		// return
		return nil, err
	}

	return userDb, nil
}

func CheckPassword(user *models.User, pw string) (error) {
	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pw))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		fmt.Println(errf)
		return errf
	}
	return nil
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