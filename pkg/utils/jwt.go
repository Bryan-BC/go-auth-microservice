package utils

import (
	"log"
	"time"

	"github.com/Bryan-BC/go-auth-microservice/pkg/models"
	"github.com/golang-jwt/jwt"
)

type JWTWrapper struct {
	Secret          string
	Issuer          string
	ExpirationHours int
}

type jwtClaims struct {
	jwt.StandardClaims
	Id       int64
	Username string
}

func (w *JWTWrapper) GenerateToken(user models.User) (string, error) {
	claims := &jwtClaims{
		Id:       user.Id,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(w.ExpirationHours)).Unix(),
			Issuer:    w.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(w.Secret))

	if err != nil {
		log.Panicf("Error signing token, %s", err)
		return "", err
	}

	return signedToken, nil
}

func (w *JWTWrapper) ValidateToken(signedToken string) (*jwtClaims, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&jwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(w.Secret), nil
		},
	)

	if err != nil {
		log.Panicf("Error parsing token, %s", err)
		return nil, err
	}

	claims, ok := token.Claims.(*jwtClaims)

	if !ok {
		log.Panicf("Error casting token claims, %s", err)
		return nil, err
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		log.Panicf("Token expired, %s", err)
		return nil, err
	}

	return claims, nil
}
