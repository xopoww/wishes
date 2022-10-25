package auth

import "github.com/xopoww/wishes/internal/db"

func GenerateToken(user *db.User) (string, error) {
	return "my-awesome-token", nil
}

func ValidateToken(token string) bool {
	return token == "my-awesome-token"
}