package hash

import (
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

const charset = "lihaoyu" + "CQUPT2023212929"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func CreatSalt() string {
	b := make([]byte, bcrypt.MaxCost)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
func CheckHashAndPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
