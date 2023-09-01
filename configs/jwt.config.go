package configs

import "github.com/golang-jwt/jwt/v5"

var JWT_KEY = []byte("ininamanyakeyyamassehjanganlupa")

type JWTClaims struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}
