package cryptography

import (
	"EducationProject/internal/domain"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWTManager struct {
	Secret []byte
	TTL    time.Duration
	now    func() time.Time
}

type JWTClaims struct {
	UserID int64 `json:"user_id"`
	//Анонимное встраивание
	jwt.RegisteredClaims
}

func NewJWTManager(secret string, ttl time.Duration) *JWTManager {
	return &JWTManager{Secret: []byte(secret), TTL: ttl, now: time.Now}
}

func (j *JWTManager) GenerateToken(UserId int64) (domain.AuthToken, error) {
	now := j.now().UTC()
	expiresAt := now.Add(j.TTL)

	claims := JWTClaims{UserID: UserId,
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt: jwt.NewNumericDate(now),
		}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString(j.Secret)
	if err != nil {
		return domain.AuthToken{}, err
	}
	return domain.AuthToken{AccessToken: accessToken, ExpiresAt: expiresAt}, nil
}

func (j *JWTManager) ParseToken(token string) (domain.TokenClaims, error) {
	parsedToken, err := jwt.ParseWithClaims(
		token,
		&JWTClaims{},
		func(incomingToken *jwt.Token) (any, error) {
			if incomingToken.Method == nil || incomingToken.Method.Alg() != jwt.SigningMethodHS256.Alg() {
				return nil, domain.ErrNonAuthorized
			}
			return j.Secret, nil
		},
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return domain.TokenClaims{}, err
	}
	claims, ok := parsedToken.Claims.(*JWTClaims)
	if !ok || !parsedToken.Valid || claims.UserID <= 0 || claims.ExpiresAt == nil {
		return domain.TokenClaims{}, domain.ErrNonAuthorized
	}
	if claims.ExpiresAt.Time.Before(j.now()) {
		return domain.TokenClaims{}, domain.ErrNonAuthorized
	}
	return domain.TokenClaims{UserId: claims.UserID, ExpiresAt: claims.ExpiresAt.Time}, nil
}
