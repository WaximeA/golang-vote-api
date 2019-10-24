package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/WaximeA/golang-vote-api/middleware"
	"github.com/WaximeA/golang-vote-api/models"

	"github.com/gorilla/mux"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func main() {
	connString := "host=golang-vote-api_db_1 user=postgres password=secret dbname=api_vote sslmode=disable"
	db, err := gorm.Open("postgres", connString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	/*err = db.Ping()
	if err != nil {
		panic(err)
	}*/

	if !db.HasTable(&models.User{}) {
		db.AutoMigrate(&models.User{})
	}

	if !db.HasTable(&models.Vote{}) {
		db.AutoMigrate(&models.Vote{})
	}

	InitStore(dbStore{db: db})
	//InitStore(dbStore{db: db})

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}

// NewRouter is used to set all routes of the project
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
