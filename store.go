package main

import (
	"github.com/WaximeA/golang-vote-api/models"
	"github.com/jinzhu/gorm"
)

//Store structure
type Store interface {
	StoreUser(user *models.User) bool
	GetStoredUser() ([]*models.User, error)
	StoreVote(vote *models.Vote) bool
	GetStoredVote() ([]*models.Vote, error)
}

var store Store

// Store user into postgres
func (store dbStore) StoreUser(User *models.User) bool {

	err := store.db.NewRecord(User) // => returns `true` as primary key is blank

	store.db.Create(&User)

	err = store.db.NewRecord(User) // => return `false` after `user` created

	return err
}

// Get stored user
func (store dbStore) GetStoredUser() ([]*models.User, error) {

	rows, err := store.db.Raw("SELECT * from users").Rows()
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
func (store dbStore) StoreVote(Vote *models.Vote) bool {
	err := store.db.NewRecord(Vote) // => returns `true` as primary key is blank

	store.db.Create(&Vote)

	err = store.db.NewRecord(Vote) // => return `false` after `Vote` created

	return err
}

// Get stored vote
func (store dbStore) GetStoredVote() ([]*models.Vote, error) {
	rows, err := store.db.Raw("SELECT * from votes").Rows()
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

// InitStore set the store
func InitStore(s Store) {
	store = s
}

// dbStore struct
type dbStore struct {
	db *gorm.DB
}
