package repository

import (
	"context"
	"database/sql"

	"github.com/locphu95/smart_machine/backend-core/internal/domain"
)

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}
func (r *UserRepositoryImpl) Create(user *domain.User) error {
	query := `INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id`
	return r.db.QueryRow(query, user.Email, user.Password).Scan(&user.ID)
}

func (r *UserRepositoryImpl) GetByEmail(email string) (*domain.User, error) {
	query := `SELECT id, email, password FROM users WHERE email = $1`
	row := r.db.QueryRow(query, email)

	var user domain.User
	if err := row.Scan(&user.ID, &user.Email, &user.Password); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) GetUser(ctx context.Context,
	id string) *domain.User {
	if id == "1" {
		return &domain.User{
			ID:   "1",
			Name: "Phu Loc V1",
		}
	}
	return nil
}

func (r *UserRepositoryImpl) GetUserV2(ctx context.Context,
	id string) *domain.User {
	if id == "2" {
		return &domain.User{
			ID:   "1",
			Name: "Phu Loc V2",
		}
	}
	return nil
}

func (r *UserRepositoryImpl) Find(
	id string) *domain.User {
	if id == "1" {
		return &domain.User{
			ID:   "1",
			Name: "Phu Loc V2",
		}
	}
	return nil
}
func (s *UserRepositoryImpl) CreateUser(ctx context.Context,
	user *domain.User) *domain.User {
	return nil
}
func (s *UserRepositoryImpl) GetUserByEmail(ctx context.Context,
	email string) *domain.User {
	return nil
}

func (s *UserRepositoryImpl) Register(ctx context.Context, email, password string) *domain.User {
	return nil
}
