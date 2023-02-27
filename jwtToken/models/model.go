package mod

import "github.com/golang-jwt/jwt/v4"

// creating struct to read username and password from request body
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
