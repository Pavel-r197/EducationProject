package postgres

import (
	"EducationProject/internal/domain"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// TODO: добавить проверку, существует ли пользователь, запретить повторную регистрацию с одним email

func (u *UserRepository) Create(ctx context.Context, user *domain.User) error {
	model := toUserModel(*user)
	if err := u.db.WithContext(ctx).Create(&model).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return domain.ErrAlreadyExists
		}
		return fmt.Errorf("insert user: %w", err)
	}
	return nil
}

func (u *UserRepository) GetById(ctx context.Context, userID int64) (*domain.User, error) {
	var user userModel
	// TODO: проверить, существует ли запись, по аналологии с проверкой email
	if err := u.db.WithContext(ctx).Where("user_id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	domainUser := toDomainUser(user)
	return &domainUser, nil
}

func (u *UserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	// TODO: дописать
	var user userModel
	if err := u.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	domainUser := toDomainUser(user)
	return &domainUser, nil
}

//func (u *UserRepository) Update(ctx context.Context, user User) error {
//
//}
