package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/lib/pq"
)

type vote struct {
	UUID  int    `json:"uuid"`
	Title string `json:"title"`
	Desc  string `json:"desc`
}

type allVotes []*vote

var votes = allVotes{
	{
		UUID:  1,
		Title: "My first vote",
		Desc:  "This is the description of my first vote.",
	},
}

func createVote(w http.ResponseWriter, r *http.Request) {

	var newVote *vote
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "There was an issue while creating the vote.")
	}

	json.Unmarshal(reqBody, &newVote)
	store.CreateVote(newVote)
	votes = append(votes, newVote)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newVote)
}

func getVotes(w http.ResponseWriter, r *http.Request) {
	store.GetVotes()
}
