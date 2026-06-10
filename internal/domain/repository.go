package domain

import "context"

type TaskRepository interface {
	Create(ctx context.Context, task *Task) error
	GetById(ctx context.Context, id, userId int64) (*Task, error)
	Update(ctx context.Context, task Task) error
	Delete(ctx context.Context, id, userId int64) error
}

// TODO: добавить функцию удаления и изменения пользователя (delete, update)

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetById(ctx context.Context, id int64) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	Update(ctx context.Context, user User) error
	//Delete()
}
