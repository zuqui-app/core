package repo

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"zuqui-core/internal/domain"
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
	return domain.User{}, nil
}

func (r *userRepo) GetUserByEmail(email string) (domain.User, error) {
	return domain.User{}, nil
}

func (r *userRepo) CreateUser(user domain.User) (domain.User, error) {
	return domain.User{}, nil
}

func (r *userRepo) UpdateUser(user domain.User) (domain.User, error) {
	return domain.User{}, nil
}

func (r *userRepo) DeleteUser(id string) error {
	return nil
}
