package main

import (
	"fmt"

	"github.com/go-jose/go-jose/v3"
)

func main() {
	// define your payload
	payload := []byte(`{"sub":"1234567890","name":"John Doe","iat":1516239022}`)

	// create a new signer using a symmetric key
	key := []byte("mysecret")
	signer, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.HS256, Key: key}, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Signer is:", signer)
	// create a JWE encrypter
	encrypter, err := jose.NewEncrypter(jose.A128GCM, jose.Recipient{Algorithm: jose.DIRECT, Key: key}, nil)
	if err != nil {
		panic(err)
	}

	// create a JWE token
	jwe, err := encrypter.Encrypt(payload)
	if err != nil {
		panic(err)
	}

	// serialize the JWE token to a string
	jweStr, err := jwe.CompactSerialize()
	if err != nil {
		panic(err)
	}

	fmt.Println(jweStr)
}
