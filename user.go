package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newUser *user
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "There is an issue with the user creation")
	}

	json.Unmarshal(reqBody, &newUser)
	store.CreateUser(newUser)
	users = append(users, newUser)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newUser)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID := mux.Vars(r)["id"]

	for _, singleUser := range users {
		if strconv.Itoa(singleUser.UUID) == userID {
			json.NewEncoder(w).Encode(singleUser)
		}
	}
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID := mux.Vars(r)["id"]
	var updatedUser user

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the user title and description only in order to update")
	}
	json.Unmarshal(reqBody, &updatedUser)

	for i, singleUser := range users {
		if strconv.Itoa(singleUser.UUID) == userID {
			if updatedUser.FirstName != "" {
				singleUser.FirstName = updatedUser.FirstName
			}
			if updatedUser.LastName != "" {
				singleUser.LastName = updatedUser.LastName
			}
			if updatedUser.Email != "" {
				singleUser.Email = updatedUser.Email
			}
			if updatedUser.Password != "" {
				singleUser.Password = updatedUser.Password
			}
			users = append(users[:i], singleUser)
			json.NewEncoder(w).Encode(singleUser)
		}
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID := mux.Vars(r)["id"]

	for i, singleUser := range users {
		if strconv.Itoa(singleUser.UUID) == userID {
			users = append(users[:i], users[i+1:]...)
			fmt.Fprintf(w, "The user with ID %v has been deleted successfully", userID)
		}
	}
}
