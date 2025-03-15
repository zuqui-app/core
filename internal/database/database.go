package database

import "github.com/jackc/pgx/v5"

type UserService struct {
	db *Database
}

func (s *UserService) GetUser() {}

type Database struct {
	client *pgx.Conn

	User *UserService
}

func New(client *pgx.Conn) *Database {
	db := &Database{
		client: client,
	}

	db.User = &UserService{db}

	return db
}
