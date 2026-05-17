package domain

import "time"

type User struct {
	Id           int64
	FirstName    string
	LastName     string
	BirthDate    time.Time
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type AuthToken struct {
	AccessToken string
	ExpiresAt   time.Time
}

type TokenClaims struct {
	UserId    int64
	ExpiresAt time.Time
}
