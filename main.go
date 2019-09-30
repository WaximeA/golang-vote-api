package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

// ID (int), UUID (string), AccessLevel (int), FirstName (string), LastName (string), Email (string), Password (string), DateOfBirth (time.Time), CreatedAt (time.Time), UpdatedAt (time.Time), DeletedAt (*time.Time)
type user struct {
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

type allUsers []*user

var users = allUsers{
	{
		UUID:      1,
		FirstName: "Waxime",
		LastName:  "AVELINE",
		Email:     "aveline.maxime@gmail.com",
		Password:  "pass",
	},
}

func main() {
	connString := "host=db user=postgres password=secret dbname=api_vote sslmode=disable"
	db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	InitStore(&dbStore{db: db})

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/users", createUser).Methods("POST")
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.Use(LoginMiddleware)
	return router
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func createUser(w http.ResponseWriter, r *http.Request) {
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
	json.NewEncoder(w).Encode(users)
}
