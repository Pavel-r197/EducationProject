package postgres

import (
	"EducationProject/internal/domain"
	"time"
)

type userModel struct {
	Id           int64     `gorm:"column:id;primaryKey"`
	FirstName    string    `gorm:"column:first_name;type:varchar(255);not null"`
	LastName     string    `gorm:"column:last_name;type:varchar(255);not null"`
	BirthDate    time.Time `gorm:"column:birth_date"`
	Email        string    `gorm:"column:email;type:varchar(255);not null;unique"`
	PasswordHash string    `gorm:"column:password_hash;type:text;not null"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}

func (userModel) TableName() string {
	return "users"
}

type taskModel struct {
	Id          int64     `gorm:"column:id;primaryKey"`
	Title       string    `gorm:"column:title;type:varchar(255);not null"`
	Description string    `gorm:"column:description;type:varchar(255);not null"`
	UserId      int64     `gorm:"column:user_id"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func (taskModel) TableName() string {
	return "tasks"
}

func toUserModel(user domain.User) userModel {
	return userModel{Id: user.Id, FirstName: user.FirstName, LastName: user.LastName, BirthDate: user.BirthDate, Email: user.Email, PasswordHash: user.PasswordHash, CreatedAt: user.CreatedAt, UpdatedAt: user.UpdatedAt}
}

func toDomainUser(user userModel) domain.User {
	return domain.User{Id: user.Id, FirstName: user.FirstName, LastName: user.LastName, BirthDate: user.BirthDate, Email: user.Email, PasswordHash: user.PasswordHash, CreatedAt: user.CreatedAt, UpdatedAt: user.UpdatedAt}
}

func toTaskModel(task domain.Task) taskModel {
	return taskModel{Id: task.Id, Title: task.Title, Description: task.Description, UserId: task.UserId, CreatedAt: task.CreatedAt, UpdatedAt: task.UpdatedAt}
}

func toDomainTask(task taskModel) domain.Task {
	return domain.Task{Id: task.Id, Title: task.Title, Description: task.Description, UserId: task.UserId, CreatedAt: task.CreatedAt, UpdatedAt: task.UpdatedAt}
}
