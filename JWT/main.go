package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

//Basic of JWT
// iat -> issued at
func main() {

	//Encode
	claim := jwt.MapClaims{
		"subject": "1234567890",
		"name":    "John Doe",
		"iat":     time.Now().Unix(),
		"expires": time.Now().Add(time.Hour * 24).Unix(),
	}
	fmt.Println("Claims is:", claim)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	fmt.Println("sdfbfjsdbfjbsbd:", *jwt.SigningMethodHS256)
	fmt.Println("Token is:", *token)
	secretKey := []byte("Hello World")
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println("Error is", err)
	}
	fmt.Println("Token string is:", tokenString)

	//Decode the token String
	parsedToken, err := jwt.ParseWithClaims(tokenString, claim, func(token *jwt.Token) (interface{}, error) {

		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("Hello World"), nil
	})
	fmt.Println()
	fmt.Println("parsed token is :", *parsedToken)
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	fmt.Println("")
	fmt.Println("Claims is:", claims)
	if err != nil {
		fmt.Println("Error after Parsing is:", err)
	} else if ok && parsedToken.Valid {
		fmt.Println(claims["subject"], claims["name"])
	} else {
		fmt.Println("Invalid Token")
	}
}
