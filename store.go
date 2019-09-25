package main

import "database/sql"

type Store interface {
	CreateUser(user *user) error
	GetUser() ([]*user, error)
}

var store Store

func (store *dbStore) CreateUser(User *user) error {
	_, err := store.db.Query("INSERT INTO users (Firstname, Lastname, Email) VALUES ($1, $2, $3)", User.FirstName, User.LastName, User.Email)
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

func InitStore(s Store) {
	store = s
}

type dbStore struct {
	db *sql.DB
}
