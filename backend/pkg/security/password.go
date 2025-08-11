package security

import (
	"crypto/rand"
	"encoding/base64"
	"os"

	"golang.org/x/crypto/bcrypt"
)

// ランダムSalt生成（lengthはバイト数）
func GenerateSalt(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}

// bcrypt + pepper + saltでハッシュ化
func HashPassword(password, salt string) (string, error) {
	pepper := os.Getenv("PEPPER")
	peppered := password + salt + pepper

	hash, err := bcrypt.GenerateFromPassword([]byte(peppered), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// ハッシュ検証
func VerifyPassword(password, salt, hash string) bool {
	pepper := os.Getenv("PEPPER")
	peppered := password + salt + pepper

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(peppered))
	return err == nil
}
