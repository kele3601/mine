package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"mine/internal/db/model"
	"os"
)

var secretKey = []byte(os.Getenv("SECRET_KEY"))

func GenerateJwtToken(user model.User) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, JwtUserClaims{
		UserID: user.ID,
	})
	return claims.SignedString(secretKey)
}

func ParseJwtToken(tokenString string) (*JwtUserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtUserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	// check claims
	if nil != err {
		return nil, err
	} else {
		// token 过期
		if !token.Valid {
			return nil, fmt.Errorf("token已失效")
		}
		//
		if claims, ok := token.Claims.(*JwtUserClaims); !ok {
			return nil, fmt.Errorf("非法token")
		} else {
			return claims, nil
		}
	}
}

type JwtUserClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}
