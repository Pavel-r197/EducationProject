package cryptography

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type ByCryptHasher struct {
	cost int
}

func NewHasher(cost int) *ByCryptHasher {
	return &ByCryptHasher{cost: cost}
}

func (h *ByCryptHasher) PasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), h.cost)
	return string(bytes), err
}

func (h *ByCryptHasher) Compare(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err == nil {
		return nil
	}
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return errors.New("Пароли не совпадают")
	}
	return err
}
