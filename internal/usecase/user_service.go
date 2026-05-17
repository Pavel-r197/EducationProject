package usecase

import (
	"EducationProject/internal/domain"
	"context"
)

type UserService interface {
	SignUp(ctx context.Context, input SignUpInput) error
	Login(ctx context.Context, input LoginInput) (domain.AuthToken, error)
}
