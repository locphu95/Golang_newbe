package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/locphu95/smart_machine/backend-core/internal/domain"
	"github.com/locphu95/smart_machine/backend-core/internal/domain/repository"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo  repository.UserReponsitory
	redis *redis.Client
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
	user *domain.User) (domain.User, error) {

	_user := s.repo.CreateUser(ctx, user)

	if _user == nil {
		return domain.User{}, domain.ErrUserNotFound
	}

	return *_user, nil
}
func (s *UserService) Register(ctx context.Context, email, password string) (*domain.User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Email:    email,
		Password: string(hashed),
	}

	if err := s.repo.CreateUser(ctx, user); err != nil {
		return nil, domain.ErrUserNotFound
	}
	return user, nil
}
func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	// 1. Check Redis
	val, err := s.redis.Get(ctx, "user:"+email).Result()
	if err == nil {
		var user domain.User
		if jsonErr := json.Unmarshal([]byte(val), &user); jsonErr == nil {
			return &user, nil
		}
	}
	// 2. Nếu cache miss → query DB
	user := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	// 3. Ghi lại vào Redis
	data, _ := json.Marshal(user)
	s.redis.Set(ctx, "user:"+email, data, 10*time.Minute)
	return user, nil
}
