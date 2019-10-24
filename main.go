package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/WaximeA/golang-vote-api/middleware"
	"github.com/WaximeA/golang-vote-api/models"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

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
	router.HandleFunc("/login", middleware.Login).Methods("POST")
	router.HandleFunc("/users", models.CreateUser).Methods("POST")
	router.HandleFunc("/users", models.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", models.GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", models.UpdateUser).Methods("PATCH")
	router.HandleFunc("/users/{id}", models.DeleteUser).Methods("DELETE")
	router.HandleFunc("/votes", models.CreateVote).Methods("POST")
	router.HandleFunc("/votes", models.GetVotes).Methods("GET")
	router.HandleFunc("/votes/{id}", models.GetVote).Methods("GET")
	router.Use(middleware.LoginMiddleware)
	return router
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}
