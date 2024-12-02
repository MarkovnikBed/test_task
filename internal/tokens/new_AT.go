package tokens

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GetNewAccessToken(ip string, id string, issueAt int64) (string, error) {
	claims := Claims{
		Ip: ip,
		Id: id,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  issueAt,
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
		},
	}
	tok := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenStr, err := tok.SignedString(Key)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}
