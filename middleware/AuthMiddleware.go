// auth/auth.go

package middleware

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go-dp.abdanhafidz.com/config"
	"go-dp.abdanhafidz.com/models"
	"golang.org/x/crypto/bcrypt"
)

// Define a secret key for signing the JWT token
var salt = config.Salt
var secretKey = []byte(salt)

// GenerateToken generates a JWT token for the given user
func GenerateToken(user *models.Account) (string, error) {

	// Create a new token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.IDAccount
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// VerifyPassword verifies if the provided password matches the hashed password
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

type CustomClaims struct {
	jwt.StandardClaims
	IDUser int `json:"id"`
}

func VerifyToken(bearer_token string) (int, string, error) {
	// fmt.Println(bearer_token)
	token, err := jwt.ParseWithClaims(bearer_token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return 0, "invalid-token", err
	}
	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return 0, "invalid-token", err
	} else if claims.StandardClaims.ExpiresAt != 0 && claims.ExpiresAt < time.Now().Unix() {
		return 0, "expired", err
	} else if !ok && token.Valid {
		return 0, "invalid-token", err
	}

	return claims.IDUser, "valid", err
}

func AuthUser(c *gin.Context) {
	var currAccData models.AccountData
	if c.Request.Header["Auth-Bearer-Token"] != nil {
		token := c.Request.Header["Auth-Bearer-Token"]
		currAccData.IdUser, currAccData.VerifyStatus, currAccData.ErrVerif = VerifyToken(token[0])
		// fmt.Println("Verify Status :", currAccData.verifyStatus)
		if currAccData.VerifyStatus == "invalid-token" || currAccData.VerifyStatus == "expired" {
			currAccData.IdUser = 0
			message := "Your session is expired, Please re-Login!"
			SendJSON401(c, &currAccData.VerifyStatus, &message)
			c.Abort()
			return
		}
	} else {
		currAccData.IdUser = 0
		currAccData.VerifyStatus = "no-token"
		currAccData.ErrVerif = nil
		message := "You have to Login First!"
		SendJSON401(c, &currAccData.VerifyStatus, &message)
		c.Abort()
		return
	}

	c.Set("accountData", currAccData)
	c.Next()
}
