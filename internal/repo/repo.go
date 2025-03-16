package repo

import "github.com/jackc/pgx/v5/pgxpool"

type Repo struct {
	User UserRepo
}

func New(db *pgxpool.Pool) *Repo {
	return &Repo{
		User: newUserRepo(db),
	}
}
