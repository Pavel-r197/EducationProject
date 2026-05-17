package domain

type PasswordHasher interface {
	PasswordHash(password string) (string, error)
	Compare(hash, password string) error
}

type TokenManager interface {
	GenerateToken(userID int64) (AuthToken, error)
	ParseToken(token string) (TokenClaims, error)
}
