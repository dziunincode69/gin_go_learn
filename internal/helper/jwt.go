package helper

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type MyAppClaims struct {
	jwt.RegisteredClaims
	ID    int
	Email string
}

func NewSign(Data map[string]any) (string, error) {
	key := os.Getenv("JWT_KEY")
	ID := Data["id"].(int)
	Email := Data["email"].(string)
	claims := MyAppClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "ManestySimpleApp",
		},
		ID:    ID,
		Email: Email,
	}
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)
	jwttoken, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return jwttoken, nil
}
func ParseToken(access_token string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(access_token, &MyAppClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
func CheckJWT(tokenjwt *jwt.Token) (*MyAppClaims, error) {
	claim := tokenjwt.Claims.(*MyAppClaims)
	if time.Now().Unix() > claim.ExpiresAt.Unix() {
		return nil, errors.New("JWT Has Expired")
	}
	return claim, nil
}
