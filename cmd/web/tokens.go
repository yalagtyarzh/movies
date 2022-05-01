package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/pascaldekloe/jwt"
	"github.com/yalagtyarzh/movies/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

var validUser = models.User{
	ID:       1,
	Email:    "lempionnewliving@gmail.com",
	Password: "$2a$12$az7EJ.t4QwX8Xg6I8OXKVO/EMo7O1Euxe/6xh3XzhRd.czUB8WVri",
}

type Credentials struct {
	Username string `json:"email"`
	Password string `json:"password"`
}

func (app *application) Signin(w http.ResponseWriter, r *http.Request) {
	var creds Credentials

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		app.errorJSON(w, errors.New("unauthorized"))
		return
	}

	hashedPassword := validUser.Password

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(creds.Password))
	if err != nil {
		app.errorJSON(w, errors.New("unauthorized"))
		return
	}

	var claims jwt.Claims
	claims.Subject = fmt.Sprint(validUser.ID)
	claims.Issued = jwt.NewNumericTime(time.Now())
	claims.NotBefore = jwt.NewNumericTime(time.Now())
	claims.Expires = jwt.NewNumericTime(time.Now().Add(24 * time.Hour))
	claims.Issuer = "AlisterAgraelAzimuth.com"
	claims.Audiences = []string{"AlisterAgraelAzimuth.com"}

	jwtBytes, err := claims.HMACSign(jwt.HS256, []byte(app.config.jwt.secret))
	if err != nil {
		app.errorJSON(w, errors.New("error signing"))
		return
	}

	app.writeJSON(w, http.StatusOK, string(jwtBytes), "response")
}
