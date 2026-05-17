package domain

import "errors"

//TODO: добавить ошибки из других участков кода

var (
	ErrNotFound      = errors.New("Запись не найдена")
	ErrInvalidInput  = errors.New("Невалидный input")
	ErrInvalidEmail  = errors.New("Невалидный email")
	ErrEmptyPassword = errors.New("Пароль не может быть пустым")
	ErrMinPassLength = errors.New("Пароль не может быть меньше 8 символов")
	ErrMaxPassLength = errors.New("Пароль не может быть больше 20 символов")
)
