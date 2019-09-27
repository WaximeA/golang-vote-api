package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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
}

func extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("bad header value given")
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("incorrectly formatted authorization header")
	}

	return jwtToken[1], nil
}

func parseToken(jwtToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, OK := token.Method.(*jwt.SigningMethodHMAC); !OK {
			return nil, errors.New("bad signed method received")
		}
		return []byte("supersaucysecret"), nil
	})

	if err != nil {
		return nil, errors.New("bad jwt token")
	}

	return token, nil
}

func jwtTokenCheck(w http.ResponseWriter, r *http.Request) bool {
	if r.RequestURI == "/login" {
		return true
	}
	jwtToken, err := extractBearerToken(r.Header.Get("Authorization"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(UnsignedResponse{
			Message: err.Error(),
		})
		return false
	}

	token, err := parseToken(jwtToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(UnsignedResponse{
			Message: "bad jwt token",
		})
		return false
	}

	_, OK := token.Claims.(jwt.MapClaims)
	if !OK {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(UnsignedResponse{
			Message: "unable to parse claims",
		})
		return false
	}
	return true
}

func LoginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		OK := jwtTokenCheck(w, r)
		if OK {
			next.ServeHTTP(w, r)
		}
	})
}
