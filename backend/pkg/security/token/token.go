package token

import (
	"fmt"
	"inoUwU/pinu/app/domain/port"
	"inoUwU/pinu/app/domain/user"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTMaker struct {
	secret []byte
}

func NewJwtMaker(secretKey string) (*JWTMaker, error) {
	return &JWTMaker{
		secret: []byte(secretKey),
	}, nil
}

// GenerateToken ユーザー情報からTokenを作成する（port.TokenMaker を実装）
func (maker *JWTMaker) GenerateToken(u *user.User, duration time.Duration) (string, *port.Claims, error) {
	claims, err := NewUserClaims(string(u.UserID), string(u.LoginID), u.IsAdmin, duration)

	if err != nil {
		return "", nil, err
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := t.SignedString(maker.secret)
	if err != nil {
		return "", nil, fmt.Errorf("error signing token: %v", err)
	}

	portClaims := &port.Claims{
		ID:        claims.RegisteredClaims.ID,
		UserID:    claims.UserId,
		LoginID:   claims.LoginID,
		IsAdmin:   claims.IsAdmin,
		IssuedAt:  claims.RegisteredClaims.IssuedAt.Time,
		ExpiresAt: claims.RegisteredClaims.ExpiresAt.Time,
	}

	return tokenStr, portClaims, nil
}

// VerifyToken JwtTokenを検証する（port.TokenMaker を実装）
func (maker *JWTMaker) VerifyToken(tokenString string) (*port.Claims, error) {
	parsed, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return maker.secret, nil
	})
	if err != nil {
		return nil, err
	}
	userClaims, ok := parsed.Claims.(*UserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return &port.Claims{
		ID:        userClaims.RegisteredClaims.ID,
		UserID:    userClaims.UserId,
		LoginID:   userClaims.LoginID,
		IsAdmin:   userClaims.IsAdmin,
		IssuedAt:  userClaims.RegisteredClaims.IssuedAt.Time,
		ExpiresAt: userClaims.RegisteredClaims.ExpiresAt.Time,
	}, nil
}
