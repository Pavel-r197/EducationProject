package stdsql

import (
	"EducationProject/internal/domain"
	"context"
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) Create(ctx context.Context, user *domain.User) error {
	const query = `
    INSERT INTO users (first_name, last_name, birth_date, email, password_hash)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING id, created_at, updated_at;`
	err := u.db.QueryRowContext(ctx, query, user.FirstName, user.LastName, user.BirthDate, user.Email, user.PasswordHash).Scan(&user.Id, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) GetById(ctx context.Context, id int64) (*domain.User, error) {
	const query = `
    SELECT id
    FROM users
    WHERE id = $1;`
	var user domain.User
	err := u.db.QueryRowContext(ctx, query, id).Scan(&user.Id, &user.LastName, &user.BirthDate, &user.Email, user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return &domain.User{}, err
	}
	return &user, nil
}

func (u *UserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	const query = `
    SELECT email
    FROM users
    WHERE email = $1;`
	var user domain.User
	err := u.db.QueryRowContext(ctx, query, email).Scan(&user.Id, &user.LastName, &user.BirthDate, &user.Email, user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return &domain.User{}, err
	}
	return &user, nil
}

//func (u *UserRepository) Update(ctx context.Context, user User) error {
//
//}
