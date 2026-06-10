package postgres

import (
	"EducationProject/internal/domain"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (t *TaskRepository) Create(ctx context.Context, task *domain.Task) error {
	model := toTaskModel(*task)
	if err := t.db.WithContext(ctx).Create(&model).Error; err != nil {
		return err
	}
	*task = toDomainTask(model)
	return nil
}

func (t *TaskRepository) GetById(ctx context.Context, id, UserID int64) (*domain.Task, error) {
	// TODO: дописать
	// Получаем id задачи и возвращаем задачу по id, если не найдена, возвращаем ошибку
	var task taskModel
	if err := t.db.WithContext(ctx).Where("id= ? and user_id= ?", id, UserID).First(&task).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("select task: %w", err)
	}
	domainTask := toDomainTask(task)
	return &domainTask, nil

}

func (t *TaskRepository) Update(ctx context.Context, task domain.Task) error {
	var model taskModel
	if err := t.db.WithContext(ctx).Where("id= ? and user_id= ?", task.Id, task.UserId).First(&model).Error; err != nil {
		return domain.ErrNotFound
	}
	model.Title = task.Title
	model.Description = task.Description
	if err := t.db.WithContext(ctx).Save(&model).Error; err != nil {
		return err
	}
	return nil
}

func (t *TaskRepository) Delete(ctx context.Context, id, UserId int64) error {
	var task taskModel
	res := t.db.WithContext(ctx).Where("id= ? and user_id= ?", id, UserId).Delete(&task)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return domain.ErrNotFound
	}
	return nil
}
