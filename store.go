package main

import (
	"database/sql"
	"github.com/WaximeA/golang-vote-api/models"
)

type Store interface {
	CreateUser(user *models.User) error
	GetUser() ([]*models.User, error)
	CreateVote(vote *models.Vote) error
	GetVotes() ([]*models.Vote, error)
}

var store Store

// Create user into postgres
func (store *dbStore) CreateUser(User *models.User) error {
	_, err := store.db.Query("INSERT INTO users (firstname, lastname, email) VALUES ($1, $2, $3)", User.FirstName, User.LastName, User.Email)
	return err
}

func (store *dbStore) GetUser() ([]*models.User, error) {
	rows, err := store.db.Query("SELECT * from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*models.User{}
	for rows.Next() {
		user := &models.User{}
		if err := rows.Scan(&user.FirstName, &user.LastName, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (store *dbStore) CreateVote(Vote *models.Vote) error {
	_, err := store.db.Query("INSERT INTO votes ( uuid, title, description) VALUES ($1, $2, $3) ", Vote.UUID, Vote.Title, Vote.Desc)
	return err
}

func (store *dbStore) GetVotes() ([]*models.Vote, error) {
	rows, err := store.db.Query("SELECT * from votes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	votes := []*models.Vote{}

	for rows.Next() {
		vote := &models.Vote{}
		if err := rows.Scan(&vote.UUID, &vote.Title, &vote.Desc); err != nil {
			return nil, err
		}
		votes = append(votes, vote)
	}

	return votes, nil
}

func InitStore(s Store) {
	store = s
}

type dbStore struct {
	db *sql.DB
}
