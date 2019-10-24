package models

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
)

type Vote struct {
	UUID  int    `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

type allVotes []*Vote

var votes = allVotes{
	{
		UUID:  1,
		Title: "My first vote",
		Desc:  "This is the description of my first vote.",
	},
}

// Create vote
func CreateVote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newVote *Vote
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "There was an issue while creating the vote.")
	}

	json.Unmarshal(reqBody, &newVote)
	//main.CreateVote(newVote)
	votes = append(votes, newVote)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newVote)
}

// Get all votes
func GetVotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(votes)
}

// Get specific vote
func GetVote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	voteID := mux.Vars(r)["id"]

	for _, singleVote := range votes {
		if strconv.Itoa(singleVote.UUID) == voteID {
			json.NewEncoder(w).Encode(singleVote)
		}
	}
}
