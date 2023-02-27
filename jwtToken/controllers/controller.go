package cont

import (
	"encoding/json"
	"fmt"
	mod "jwt-token/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var users = []mod.Credentials{
	{Username: "user1", Password: "password1"},
	{Username: "user2", Password: "password2"},
}
var SecretKey = []byte("my_secret_key")

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user mod.Credentials
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	users = append(users, user)
	json.NewEncoder(w).Encode(users)
}

// In this function we create the token and check for password validity
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var cred mod.Credentials
	// Get the JSON body and decode into credentials
	err := json.NewDecoder(r.Body).Decode(&cred)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(cred)

	var expectPass string

	for _, val := range users {
		if cred.Username == val.Username {
			expectPass = val.Password
		}
	}

	bs, err := bcrypt.GenerateFromPassword([]byte(expectPass), 2)
	loginPass := cred.Password

	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPass))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &mod.Claims{
		Username: cred.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println("Token String is :", tokenString)
	cookie := &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: time.Now().Add(time.Hour * 24),
	}
	http.SetCookie(w, cookie)

	w.Write([]byte(fmt.Sprint("Token is:", tokenString)))
}

// In this function we will check for the validity of the token if valid then dsisplay the welcome message
func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	//before this middleware is running to check if the token is valid or not
	w.Write([]byte(fmt.Sprintf("Welcome User")))
}

func RefreshHandler(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tknStr := c.Value
	claims := &mod.Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = jwt.NewNumericDate(expirationTime)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Expires: time.Now(),
	})

	w.Write([]byte("Logged out sucessfully"))
}

func MiddleWare(originHandler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		fmt.Println("Cokkie is :", c)

		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			w.WriteHeader(http.StatusBadRequest)
			return
		}

		tokenString := c.Value
		claims := &mod.Claims{}

		parsedToken, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return SecretKey, nil
		})
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !parsedToken.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		originHandler.ServeHTTP(w, r)

	})
}
