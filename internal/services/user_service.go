package service

import (
	"context"

	"github.com/locphu95/smart_machine/backend-core/internal/domain"
	"github.com/locphu95/smart_machine/backend-core/internal/domain/repository"
)

type UserService struct {
	repo repository.UserReponsitory
}

func NewUserService(repo repository.UserReponsitory) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetUser(ctx context.Context,
	id string) (*domain.User, error) {

	user := s.repo.GetUser(ctx, id)

	if user == nil {
		return nil, domain.ErrUserNotFound
	}

	return user, nil
}

func (s *UserService) GetUserV2(ctx context.Context,
	id string) (domain.User, error) {
	user := s.repo.Find(id)

	if user == nil {
		return domain.User{}, domain.ErrUserNotFound
	}

	return *user, nil
}

func (s *UserService) CreateUser(ctx context.Context,
	user domain.User) (domain.User, error) {

	_user := s.repo.CreateUser(ctx, user)

	if _user == nil {
		return domain.User{}, domain.ErrUserNotFound
	}

	return *_user, nil
}
