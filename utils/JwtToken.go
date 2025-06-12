package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// userName编码生成token
func CreateToken(userName string, secret string) (string, error) {
	at := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"userName": userName,
		"exp":      time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

// token解码获取userName
func ParseToken(token string, secret string) (string, error) {
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}
	return claim.Claims.(jwt.MapClaims)["userName"].(string), nil
}
