package usecase

import (
	"EducationProject/internal/domain"
	"context"
)

type TaskService interface {
	Create(ctx context.Context, input TaskInput) (domain.Task, error)
	GetById(ctx context.Context, id int64) (domain.Task, error)
	Update(ctx context.Context, input UpdateTask) error
	Delete(ctx context.Context, id int64) error
}
