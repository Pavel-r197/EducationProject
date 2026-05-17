package usecase

import (
	"EducationProject/internal/domain"
	"context"
	"errors"
)

const (
	DefaultTaskPage  = 1
	DefaultTaskLimit = 20
	MaxTaskLimit
)

type TaskInput struct {
	Title       string
	Description string
}

type UpdateTask struct {
	Id          int64
	Title       string
	Description string
}

type TaskUseCase struct {
	Tasks domain.TaskRepository
}

func NewTaskUseCase(n domain.TaskRepository) TaskService {
	return &TaskUseCase{Tasks: n}
}

func (n *TaskUseCase) Create(ctx context.Context, input TaskInput) (domain.Task, error) {
	// TODO: добавить проверку id пользователя
	if len(input.Title) < 3 || len(input.Description) < 3 {
		return domain.Task{}, domain.ErrInvalidInput
		// TO-DO: подключить domain errors.go для вывода ошибок
	}
	task := domain.Task{Title: input.Title, Description: input.Description, UserId: 1}
	if err := n.Tasks.Create(ctx, &task); err != nil {
		return domain.Task{}, errors.New("Ошибка создания задачи")
	}

	return task, nil

}

func (n *TaskUseCase) GetById(ctx context.Context, id int64) (domain.Task, error) {
	if id < 1 {
		return domain.Task{}, errors.New("ID не может быть меньше 1")
	}
	task, err := n.Tasks.GetById(ctx, id, 1)
	if err != nil {
		return domain.Task{}, errors.New("В GetById что то пошло не так")
	}
	return *task, nil
}

func (n *TaskUseCase) Update(ctx context.Context, input UpdateTask) error {
	if len(input.Title) < 3 || len(input.Description) < 3 {
		return errors.New("Invalid input")
		// TO-DO: подключить domain errors.go для вывода ошибок
	}
	task := domain.Task{Id: input.Id, Title: input.Title, Description: input.Description, UserId: 1}
	if err := n.Tasks.Update(ctx, task); err != nil {
		return err
	}
	return nil
}

func (n *TaskUseCase) Delete(ctx context.Context, id int64) error {
	if id < 1 {
		return errors.New("ID не может быть меньше 1")
	}
	err := n.Tasks.Delete(ctx, id, 1)
	if err != nil {
		return err
	}
	return nil
}
