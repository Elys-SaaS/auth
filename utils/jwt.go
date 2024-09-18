package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTType string

const (
	AccessToken  JWTType = "access"
	RefreshToken JWTType = "refresh"
)

var JWTAccessSecret = []byte("!!SECRET!!")
var JWTRefreshSecret = []byte("!!SECRET!!")

func GenerateJWT(id uint, jwtType JWTType) string {
	secret := JWTAccessSecret
	exp := time.Now().Add(time.Hour * 48).Unix()
	if jwtType == RefreshToken {
		exp = time.Now().Add(time.Hour * 24 * 7).Unix()
		secret = JWTRefreshSecret
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = exp
	t, _ := token.SignedString(secret)

	return t
}

func VerifyJWT(t string, jwtType JWTType) (*jwt.MapClaims, error) {
	secret := JWTAccessSecret
	if jwtType == RefreshToken {
		secret = JWTRefreshSecret
	}
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}
	return &claims, nil
}
