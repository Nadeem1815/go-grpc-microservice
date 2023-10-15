package utils

import (
	"errors"
	"time"

	"github.com/Nadeem1815/go-grpc-auth-svc/pkg/models"
	"github.com/golang-jwt/jwt"
)

// we generate and validate JSON Web Tokens based on a secret key that we defined earlier in the dev.env file
type JwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

type jwtClaims struct {
	jwt.StandardClaims
	Id    int64
	Email string
}

func (w *JwtWrapper) GenerateToken(user models.User) (signedToken string, err error) {
	claims := &jwtClaims{
		Id:    user.Id,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(w.ExpirationHours)).Unix(),
			Issuer:    w.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err = token.SignedString([]byte(w.SecretKey))

	if err != nil {
		return "", nil

	}
	return signedToken, nil
}

func (w *JwtWrapper) ValidateToken(signedToken string) (claims *jwtClaims, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&jwtClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(w.SecretKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*jwtClaims)

	if !ok {
		return nil, errors.New("couldn't parse claims")

	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("Jwt is expaired")
	}
	return claims, nil
}
