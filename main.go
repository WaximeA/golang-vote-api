package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
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
	router.HandleFunc("/users/{id}", getUser).Methods("GET")
	router.HandleFunc("/users/{id}", updateUser).Methods("PATCH")
	router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")
	router.HandleFunc("/votes", createVote).Methods("POST")
	router.HandleFunc("/votes", getVotes).Methods("GET")
	router.Use(LoginMiddleware)
	return router
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}
