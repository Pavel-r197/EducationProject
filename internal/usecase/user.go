package usecase

import (
	"EducationProject/internal/domain"
	"context"
	"errors"
	"net/mail"
	"strings"
	"time"
)

type SignUpInput struct {
	FirstName string
	LastName  string
	BirthDate time.Time
	Email     string
	Password  string
}

type LoginInput struct {
	Email    string
	Password string
}

type AuthUseCase struct {
	Users  domain.UserRepository
	hasher domain.PasswordHasher
	token  domain.TokenManager
}

func NewAuthUseCase(u domain.UserRepository, h domain.PasswordHasher, t domain.TokenManager) AuthUseCase {
	return AuthUseCase{Users: u, hasher: h, token: t}
}

func (u AuthUseCase) SignUp(ctx context.Context, input SignUpInput) error {
	ok := emailValidator(input.Email)
	if !ok {
		return domain.ErrInvalidEmail
	}
	// Удаляем пробелы по краям строки
	var firstName = strings.TrimSpace(input.FirstName)
	var lastName = strings.TrimSpace(input.LastName)
	if len(firstName) < 3 || len(lastName) < 3 {
		return errors.New("Количество символов в полях имени и фамилии должно быть больше 3")
	}

	if err := passwordValidator(input.Password); err != nil {
		return err
	}
	passwordHash, err := u.hasher.PasswordHash(input.Password)
	if err != nil {
		return errors.New("Ошибка создания хэш пароля")
	}
	var user = domain.User{LastName: lastName, FirstName: firstName, BirthDate: input.BirthDate, PasswordHash: passwordHash}
	err = u.Users.Create(ctx, &user)
	if err != nil {
		return err
	}
	return nil
}

//TODO: вернуть токен

func (u AuthUseCase) Login(ctx context.Context, input LoginInput) error {
	ok := emailValidator(input.Email)
	if !ok {
		return domain.ErrInvalidEmail
	}

	if err := passwordValidator(input.Password); err != nil {
		return err
	}
	// TODO: по аналогии с register проверить почту, пароль
	// TODO: воспользоваться методом GetByEmail, если не найдено вернуть ошибку
	// TODO: сравнить PasswordHash из структуры с input.password
	return nil
}

//func (u AuthUseCase) Logout (ctx context.Context) {
//
//}

func emailValidator(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil

}

func passwordValidator(password string) error {
	var pass = strings.TrimSpace(password)
	if pass == "" {
		return domain.ErrEmptyPassword
	}
	if len(pass) < 8 {
		return domain.ErrMinPassLength
	}
	if len(pass) > 20 {
		return domain.ErrMaxPassLength
	}
	return nil
}
