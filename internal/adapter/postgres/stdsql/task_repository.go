package stdsql

import (
	"EducationProject/internal/domain"
	"context"
	"database/sql"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (t *TaskRepository) Create(ctx context.Context, task *domain.Task) error {
	const query = `
    INSERT INTO tasks (user_id, title, description)
    VALUES ($1, $2, $3)
    RETURNING id, created_at, updated_at;`
	err := t.db.QueryRowContext(ctx, query, task.UserId, task.Title, task.Description).Scan(&task.Id, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (t *TaskRepository) GetById(ctx context.Context, id, UserID int64) (*domain.Task, error) {
	const query = `
    SELECT id, user_id, title, description, created_at, updated_at
    FROM tasks
    WHERE id = $1 AND user_id = $2;`
	var task domain.Task
	err := t.db.QueryRowContext(ctx, query, id, UserID).Scan(&task.Id, &task.UserId, &task.Title, &task.Description, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return &domain.Task{}, err
	}
	return &task, nil
}

//TODO: реализовать UPDATE

func (t *TaskRepository) Update(ctx context.Context, task domain.Task) error {
	const query = `
    UPDATE tasks
    SET title = $3, description = $4, updated_at = NOW()
    WHERE id = $1 AND user_id = $2
    RETURNING updated_at;`
	err := t.db.QueryRowContext(ctx, query, task.Id, task.UserId, task.Title, task.Description).Scan(&task.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

//TODO: реализовать DELETE

func (t *TaskRepository) Delete(ctx context.Context, id, UserId int64) error {
	const query = `
    DELETE
    FROM tasks
    WHERE id = $1 AND user_id = $2;`
	err := t.db.QueryRowContext(ctx, query, id, UserId).Scan()
	if err != nil {
		return domain.ErrTaskDelete
	}
	return nil
}
