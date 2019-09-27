package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UnsignedResponse struct {
	Message interface{} `json:"message"`
}

type SignedResponse struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}

func login(w http.ResponseWriter, r *http.Request) {
	type login struct {
		Username string `json:"username,omitempty"`
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "There is an issue with the user creation")
	}
	loginParams := login{}
	json.Unmarshal(reqBody, &loginParams)

	if loginParams.Username == "mike" || loginParams.Username == "rama" {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user": loginParams.Username,
			"nbf":  time.Date(2018, 01, 01, 12, 0, 0, 0, time.UTC).Unix(),
		})

		tokenStr, err := token.SignedString([]byte("supersaucysecret"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(UnsignedResponse{
				Message: err.Error(),
			})
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(SignedResponse{
			Token:   tokenStr,
			Message: "logged in",
		})
		return
	}
	json.NewEncoder(w).Encode(UnsignedResponse{
		Message: "bad username : " + loginParams.Username,
	})
	//c.JSON(http.StatusBadRequest, UnsignedResponse{
	//	Message: "bad username",
	//})
}
