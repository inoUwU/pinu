package token

import (
	"fmt"
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

// GenerateToken loginIdからTokenを作成します
func (maker *JWTMaker) GenerateToken(user *user.User, duration time.Duration) (string, *UserClaims, error) {
	claims, err := NewUserClaims(string(user.USER_ID), string(user.LOGIN_ID), user.IS_ADMIN, duration)

	if err != nil {
		return "", nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(maker.secret)
	if err != nil {
		return "", nil, fmt.Errorf("error signing token: %v", err)
	}

	return t, claims, nil
}

// VerifyToken JwtTokenを検証します
func (maker *JWTMaker) VerifyToken(tokenString string) (*UserClaims, error) {
	claims, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		// verify the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return maker.secret, nil
	})
	if err != nil {
		return nil, err
	}
	userClaims, ok := claims.Claims.(*UserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}
	return userClaims, nil
}
