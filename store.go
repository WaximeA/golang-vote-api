package main

import (
	"database/sql"
	"github.com/WaximeA/golang-vote-api/models"
)

type Store interface {
	StoreUser(user *models.User) error
	GetStoredUser() ([]*models.User, error)
	StoreVote(vote *models.Vote) error
	GetStoredVote() ([]*models.Vote, error)
}

var store Store

// Store user into postgres
func (store *dbStore) StoreUser(User *models.User) error {
	_, err := store.db.Query("INSERT INTO users (id, access_level, first_name, last_name, email, password, birth_date) VALUES ($1, $2, $3)", User.UUID, User.AccessLevel, User.FirstName, User.LastName, User.Email, User.Password, User.DateOfBirth)
	return err
}

// Get stored user
func (store *dbStore) GetStoredUser() ([]*models.User, error) {
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

// Store vote
func (store *dbStore) StoreVote(Vote *models.Vote) error {
	_, err := store.db.Query("INSERT INTO votes ( uuid, title, description) VALUES ($1, $2, $3) ", Vote.UUID, Vote.Title, Vote.Desc)
	return err
}

// Get stored vote
func (store *dbStore) GetStoredVote() ([]*models.Vote, error) {
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
