package repo

import "github.com/jackc/pgx/v5/pgxpool"

type Repo struct {
	user UserRepo
}

func New(db *pgxpool.Pool) *Repo {
	return &Repo{
		user: newUserRepo(db),
	}
}
