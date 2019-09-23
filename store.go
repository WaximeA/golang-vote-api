package main

import "database/sql"

type Store interface {

}

var store Store

func InitStore(s Store) {
	store = s
}

type dbStore struct {
	db *sql.DB
}