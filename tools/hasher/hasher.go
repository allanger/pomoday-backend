package hasher

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"golang.org/x/crypto/bcrypt"
)

func convertPassword(password string) []byte {
	return []byte(password)
}

func hashAndSalt(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		fmt.Print(err) //FIXME
	}
	return string(hash)
}

// Encrypt is for encrypting password and adding salt
func Encrypt(password string) string {
	return hashAndSalt(convertPassword(password))
}

// ComparePasswords compares password and hash
func ComparePasswords(hashedPassword string, plainPassword string) error {
	byteHash := []byte(hashedPassword)
	password := []byte(plainPassword)
	if err := bcrypt.CompareHashAndPassword(byteHash, password); err != nil {
		return status.Errorf(codes.PermissionDenied, "wrong password")
	}
	return nil
}
