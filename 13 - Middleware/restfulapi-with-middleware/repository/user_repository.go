package repository

import (
	"context"
	"latian_clean_arch/model/domain"
)

type UserRepository interface {
	Save(ctx context.Context, user domain.User) (domain.User, error)
	Update(ctx context.Context, user domain.User, userId uint) (domain.User, error)
	Delete(ctx context.Context, userId uint) error
	FindById(ctx context.Context, userId uint) (int, domain.User, error)
	FindByIdentifier(ctx context.Context, identifier string) (int, domain.User, error)
	FindAll(ctx context.Context) ([]domain.User, error)
	CountById(ctx context.Context, userId uint) (int, error)
	MatchingEmailAndPassword(ctx context.Context, email string, password string) (bool, error)
}
