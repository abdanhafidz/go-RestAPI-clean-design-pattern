package services

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"godp.abdanhafidz.com/config"
	"godp.abdanhafidz.com/models"
	"golang.org/x/crypto/bcrypt"
)

var salt = config.Salt
var secretKey = []byte(salt)

func GenerateToken(user *models.Account) (string, error) {

	// Create a new token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.Id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyPassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return errors.New("invalid password")
	}
	return nil
}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
