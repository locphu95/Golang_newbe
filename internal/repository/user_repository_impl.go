package repository

import (
	"context"

	"github.com/locphu95/smart_machine/backend-core/internal/domain"
)

type UserRepositoryImpl struct{}

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
	user domain.User) *domain.User {
	return nil
}
