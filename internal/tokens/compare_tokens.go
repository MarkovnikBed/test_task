package tokens

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

func CompareRTandAT(accessToken string, idRT string, issueAtRT int64) (string, error) {
	tok, err := jwt.ParseWithClaims(accessToken, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return Key, nil
	})
	if err != nil {
		return "", fmt.Errorf("ошибка при разборе токена" + err.Error())
	}
	if !tok.Valid {
		return "", fmt.Errorf("невалидный токен")
	}
	claims, ok := tok.Claims.(*Claims)
	if !ok {
		return "", fmt.Errorf("не удалось получить payload")
	}

	if claims.Id != idRT {
		return "", fmt.Errorf("id не совпадают")
	}
	if claims.IssuedAt != issueAtRT {
		return "", fmt.Errorf(
			"не совпадают времена создания",
		)
	}
	return claims.Ip, nil
}
