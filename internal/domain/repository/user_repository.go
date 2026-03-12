package repository

import (
	"context"

	"github.com/locphu95/smart_machine/backend-core/internal/domain"
)

type UserReponsitory interface {
	GetUser(ctx context.Context, id string) *domain.User
	GetUserV2(ctx context.Context, id string) *domain.User
	//Find(id string) *domain.User
	Find(id string) *domain.User
}
