package tokens

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Ip string
	Id string
	jwt.StandardClaims
}

var Key = []byte(os.Getenv("KEY"))

func GetRefreshToken(id string, ip string) (string, error) {

	claims := Claims{
		Ip: ip,
		Id: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenStr, err := token.SignedString(Key)
	if err != nil {
		return "", fmt.Errorf("не получилось создать refresh - токен авторизации")
	}
	return tokenStr, nil
}

func GetAccessToken(id string, ip string) (string, error) {

	claims := Claims{
		Ip: ip,
		Id: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenStr, err := token.SignedString(Key)
	if err != nil {
		return "", fmt.Errorf("не получилось создать access - токен авторизации")
	}
	return tokenStr, nil
}
