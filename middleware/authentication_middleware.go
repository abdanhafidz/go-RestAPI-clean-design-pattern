// auth/auth.go

package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"godp.abdanhafidz.com/config"
	"godp.abdanhafidz.com/models"
	"godp.abdanhafidz.com/utils"
)

var salt = config.Salt
var secretKey = []byte(salt)

// VerifyPassword verifies if the provided password matches the hashed password

type CustomClaims struct {
	jwt.RegisteredClaims
	UserID int `json:"id"`
}

func VerifyToken(bearer_token string) (int, string, error) {
	// fmt.Println(bearer_token)
	token, err := jwt.ParseWithClaims(bearer_token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return 0, "invalid-token", err
	}

	// Extract the claims
	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return 0, "invalid-token", err
	}
	if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
		return 0, "expired", err
	}

	return claims.UserID, "valid", err
}

func AuthUser(c *gin.Context) {
	var currAccData models.AccountData
	if c.Request.Header["Auth-Bearer-Token"] != nil {
		token := c.Request.Header["Auth-Bearer-Token"]
		currAccData.UserID, currAccData.VerifyStatus, currAccData.ErrVerif = VerifyToken(token[0])
		// fmt.Println("Verify Status :", currAccData.verifyStatus)
		if currAccData.VerifyStatus == "invalid-token" || currAccData.VerifyStatus == "expired" {
			currAccData.UserID = 0
			utils.ResponseFAIL(c, 401, models.Exception{Unauthorized: true, Message: "Your session is expired, Please re-Login!"})
			c.Abort()
			return
		} else {
			c.Set("accountData", currAccData)
			c.Next()
		}
	} else {
		currAccData.UserID = 0
		currAccData.VerifyStatus = "no-token"
		currAccData.ErrVerif = nil
		utils.ResponseFAIL(c, 401, models.Exception{Unauthorized: true, Message: "You have to login first!"})
		c.Abort()
		return
	}

}
