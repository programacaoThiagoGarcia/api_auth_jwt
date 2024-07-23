package jwtapi

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	Username string `json:"first"`
	jwt.RegisteredClaims
}

var mySigningKey = []byte(os.Getenv("API_SECRET"))

func CreateToken(username string) (string, error) {

	userClaim := UserClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 5)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaim)
	return token.SignedString(mySigningKey)

}

func VerifyToken(tokenString string) error {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	switch {
	case token.Valid:
		fmt.Println("You look nice today")
		return nil
	case errors.Is(err, jwt.ErrTokenMalformed):
		fmt.Println("That's not even a token")
		return err
	case errors.Is(err, jwt.ErrTokenSignatureInvalid):
		fmt.Println("Invalid signature")
		return err
	case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
		fmt.Println("Timing is everything")
		return err
	default:
		fmt.Println("Couldn't handle this token:", err)
		return err
	}

}
