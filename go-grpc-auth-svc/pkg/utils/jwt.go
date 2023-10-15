package utils

// import (
// 	"time"

// 	"github.com/Nadeem1815/go-grpc-auth-svc/pkg/models"
// 	"github.com/golang-jwt/jwt"
// )

// type JwtWrapper struct {
// 	SecretKey       string
// 	Issuer          string
// 	ExpirationHours int64
// }

// type jwtClaims struct {
// 	jwt.StandardClaims
// 	Id    int64
// 	Email string
// }

// func (w *JwtWrapper) GenerateToken(user models.User) (signedToken string, err error) {
// 	claims := &jwtClaims{
// 		Id:    user.Id,
// 		Email: user.Email,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(w.ExpirationHours)).Unix(),
// 			Issuer:    w.Issuer,
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt)
// }
