package models

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// ID (int), UUID (string), AccessLevel (int), FirstName (string), LastName (string), Email (string), Password (string), DateOfBirth (time.Time), CreatedAt (time.Time), UpdatedAt (time.Time), DeletedAt (*time.Time)
type User struct {
	UUID        int       `json:"id"`
	AccessLevel int       `json:"access_level"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	DateOfBirth time.Time `json:"birth_date"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type allUsers []*User

var users = allUsers{
	{
		UUID:      1,
		FirstName: "Waxime",
		LastName:  "AVELINE",
		Email:     "aveline.maxime@gmail.com",
		Password:  "pass",
	},
}

// Create user from body parameters
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newUser *User
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "There is an issue with the User creation")
	}

	json.Unmarshal(reqBody, &newUser)
	//middleware.CreateUser(newUser)
	users = append(users, newUser)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newUser)
}

// Get all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Get specific user
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID := mux.Vars(r)["id"]

	for _, singleUser := range users {
		if strconv.Itoa(singleUser.UUID) == userID {
			json.NewEncoder(w).Encode(singleUser)
		}
	}
}

// Update specific user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID := mux.Vars(r)["id"]
	var updatedUser User

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

// Delete specific user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userID := mux.Vars(r)["id"]

	for i, singleUser := range users {
		if strconv.Itoa(singleUser.UUID) == userID {
			users = append(users[:i], users[i+1:]...)
			fmt.Fprintf(w, "The user with ID %v has been deleted successfully", userID)
		}
	}
}
