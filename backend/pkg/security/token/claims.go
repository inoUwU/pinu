package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type UserClaims struct {
	UserId  string `json:"user_id"`
	LoginID string `json:"login_id"`
	IsAdmin bool   `json:"is_admin"`
	jwt.RegisteredClaims
}

func NewUserClaims(userId, loginID string, isAdmin bool, duration time.Duration) (*UserClaims, error) {
	tokenId, err := uuid.NewRandom()

	if err != nil {
		return nil, fmt.Errorf("error generating token id: %v", err)
	}

	return &UserClaims{
		UserId:  userId,
		LoginID: loginID,
		IsAdmin: isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        tokenId.String(),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	}, nil
}
