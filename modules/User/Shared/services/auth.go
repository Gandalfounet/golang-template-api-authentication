package userServices

import (
	"fmt"
	"time"
	"golang-template-api-authentication/modules/User/Shared/models"

	jwt "github.com/dgrijalva/jwt-go"

	"os"
	"github.com/joho/godotenv"
)

func GetToken(user *models.User) (string, error){
	expiresAt := time.Now().Add(time.Minute * 100000).Unix()
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
		return "", errEnv
	}
	secretJwt := os.Getenv("secretJwt")

	tokenString, error := token.SignedString([]byte(secretJwt))
	if error != nil {
		fmt.Println(error)
		return "", error
	}

	return tokenString, nil
}