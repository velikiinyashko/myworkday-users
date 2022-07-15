package authorized

import (
	"log"

	"github.com/golang-jwt/jwt"
)

func getJwtToken(login, id, key string) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"user":   login,
		"userid": id,
	})

	tokenStr, err := token.SignedString(key)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &tokenStr, nil
}
