package main

import "database/sql"

type Store interface {
	CreateUser(user *user) error
	GetUser() ([]*user, error)
	CreateVote(vote *vote) error
	GetVotes() ([]*vote, error)
}

var store Store

func (store *dbStore) CreateUser(User *user) error {
	_, err := store.db.Query("INSERT INTO users (firstname, lastname, email) VALUES ($1, $2, $3)", User.FirstName, User.LastName, User.Email)
	return err
}

func (store *dbStore) GetUser() ([]*user, error) {
	rows, err := store.db.Query("SELECT * from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*user{}
	for rows.Next() {
		user := &user{}
		if err := rows.Scan(&user.FirstName, &user.LastName, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (store *dbStore) CreateVote(Vote *vote) error {
	_, err := store.db.Query("INSERT INTO votes ( uuid, title, description) VALUES ($1, $2, $3) ", Vote.UUID, Vote.Title, Vote.Desc)
	return err
}

func (store *dbStore) GetVotes() ([]*vote, error) {
	rows, err := store.db.Query("SELECT * from votes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	votes := []*vote{}

	for rows.Next() {
		vote := &vote{}
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
