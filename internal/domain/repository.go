package domain

import "context"

// TODO: добавить функцию удаления и изменения заметок (delete, update)

type TaskRepository interface {
	Create(ctx context.Context, task *Task) error
	GetById(ctx context.Context, id, userId int64) (*Task, error)
	Update(ctx context.Context, task Task) error
	Delete(ctx context.Context, id, userId int64) error
}

// TODO: добавить функцию удаления и изменения пользователя (delete, update)

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetByEmail(ctx context.Context, email string) (*User, error)
	//Update()
	//Delete()
}
