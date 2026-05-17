package usecase

import (
	"EducationProject/internal/domain"
	"context"
)

type UserService interface {
	Register(ctx context.Context, input RegisterInput) error
	Login(ctx context.Context, input LoginInput) (domain.AuthToken, error)
}
