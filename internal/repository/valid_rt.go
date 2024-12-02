package repository

import (
	"crypto/sha256"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"medods/internal/tokens"
)

func (rep *Repository) ValidRefreshToken(refreshToken string) (IssuedAt int64, id string, ip string, err error) {

	tok, err := jwt.ParseWithClaims(refreshToken, &tokens.Claims{}, func(t *jwt.Token) (interface{}, error) {
		t.Method = jwt.SigningMethodHS512
		return tokens.Key, nil
	})

	if err != nil {
		return 0, "", "", err
	}
	if !tok.Valid {
		return 0, "", "", fmt.Errorf("невозможно распарсить RT")
	}
	cl, ok := tok.Claims.(*tokens.Claims)
	if !ok {
		return 0, "", "", fmt.Errorf("ошибка обработки токена")
	}
	row := rep.DB.QueryRow("SELECT hash from refresh_tokens WHERE id=$1", cl.Id)
	var hash string
	err = row.Scan(&hash)
	if err != nil {
		return 0, "", "", err
	}
	refreshSHA256 := sha256.Sum256([]byte(refreshToken))
	err = bcrypt.CompareHashAndPassword([]byte(hash), refreshSHA256[:])
	if err != nil {
		return 0, "", "", err
	}
	return cl.IssuedAt, cl.Id, cl.Ip, nil
}
