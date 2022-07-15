package authorized

import (
	"log"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	jwt.StandardClaims
	User string `json:"user"`
}

func GetJwtToken(login, id, key string) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		StandardClaims: jwt.StandardClaims{}, User: login,
	})

	tokenStr, err := token.SignedString([]byte(key))

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &tokenStr, nil
}
