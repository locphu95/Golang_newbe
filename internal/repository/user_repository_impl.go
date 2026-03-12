package repository

import (
	"context"

	"github.com/locphu95/smart_machine/backend-core/internal/domain"
)

type UserReponsitoryImpl struct{}

func (r *UserReponsitoryImpl) GetUser(ctx context.Context,
	id string) *domain.User {
	if id == "1" {
		return &domain.User{
			ID:   "1",
			Name: "Phu Loc V2",
		}
	}
	return nil
}

func (r *UserReponsitoryImpl) GetUserV2(ctx context.Context,
	id string) *domain.User {
	if id == "1" {
		return &domain.User{
			ID:   "1",
			Name: "Phu Loc V2",
		}
	}
	return nil
}

func (r *UserReponsitoryImpl) Find(
	id string) *domain.User {
	if id == "1" {
		return &domain.User{
			ID:   "1",
			Name: "Phu Loc V2",
		}
	}
	return nil
}
