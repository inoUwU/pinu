package pkg

import (
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"time"
)

func GenerateJWT(userID int, secretKey string) (string, error) {
	jwtSecret := []byte(secretKey)

}
