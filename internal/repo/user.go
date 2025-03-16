package repo

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/oklog/ulid/v2"

	"zuqui-core/internal/domain"
)

var (
	EmptyUser = domain.User{}
	ErrNoUser = errors.New("no user found")
)

type UserRepo interface {
	GetUserById(id string) (domain.User, error)
	GetUserByEmail(email string) (domain.User, error)
	CreateUser(user domain.User) (domain.User, error)
	UpdateUser(user domain.User) (domain.User, error)
	DeleteUser(id string) error
}

type userRepo struct {
	db *pgxpool.Pool
}

func newUserRepo(db *pgxpool.Pool) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) GetUserById(id string) (domain.User, error) {
	rows, err := r.db.Query(context.Background(), "SELECT * FROM users WHERE id = $1 LIMIT 1", id)
	if err != nil {
		return EmptyUser, nil
	}
	user, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[domain.User])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return EmptyUser, ErrNoUser
		}

		return EmptyUser, err
	}
	return user, nil
}

func (r *userRepo) GetUserByEmail(email string) (domain.User, error) {
	rows, err := r.db.Query(
		context.Background(),
		"SELECT * FROM users WHERE email = $1 LIMIT 1",
		email,
	)
	if err != nil {
		return EmptyUser, nil
	}
	user, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[domain.User])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return EmptyUser, ErrNoUser
		}

		return EmptyUser, err
	}
	return user, nil
}

func (r *userRepo) CreateUser(user domain.User) (domain.User, error) {
	rows, err := r.db.Query(
		context.Background(),
		"INSERT INTO users (id,email,username) VALUES ($1, $2, $3) RETURNING *",
		ulid.Make().String(),
		user.Email,
		user.Username,
	)
	if err != nil {
		return EmptyUser, err
	}
	user, err = pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[domain.User])
	if err != nil {
		return EmptyUser, err
	}
	return user, nil
}

func (r *userRepo) UpdateUser(user domain.User) (domain.User, error) {
	rows, err := r.db.Query(
		context.Background(),
		`
UPDATE users
SET 
    email = $2
    username = $3
WHERE
    id = $1
RETURNING *
    `,
		user.Id,
		user.Email,
		user.Username,
	)
	if err != nil {
		return EmptyUser, err
	}
	user, err = pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[domain.User])
	if err != nil {
		return EmptyUser, err
	}
	return user, nil
}

func (r *userRepo) DeleteUser(id string) error {
	_, err := r.db.Exec(context.Background(), "DELETE FROM users WHERE id = $1", id)
	return err
}
