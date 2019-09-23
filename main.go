package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type user struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Votes     int    `json:"user_votes"`
}

type allUsers []user

var users = allUsers{
	{
		ID:        1,
		FirstName: "Waxime",
		LastName:  "AVELINE",
		Email:     "aveline.maxime@gmail.com",
		Votes:     0,
	},
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/users", createUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8001", router))
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser user
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "The user was created ;)")
	} else {
		fmt.Fprintf(w, "Woupsy doopsy")
	}

	json.Unmarshal(reqBody, &newUser)
	users = append(users, newUser)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newUser)
}
