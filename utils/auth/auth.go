package auth

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"movie-app/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"github.com/xdg-go/pbkdf2"
)

type AuthJWT struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Role         string `json:"role"`
	IsRegistered bool   `json:"is_registered"`
	jwt.RegisteredClaims
}

func GenerateAuthToken(config config.Config, opts AuthJWT) (string, error) {
	claims := AuthJWT{
		ID:           opts.ID,
		Name:         opts.Name,
		Role:         opts.Role,
		IsRegistered: opts.IsRegistered,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	jwtSigner := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := jwtSigner.SignedString([]byte(config.JwtKey))
	if err != nil {
		return "", errors.WithStack(err)
	}

	return token, nil
}

func GenerateSalt() string {
	saltBytes := make([]byte, 32)
	rand.Read(saltBytes)
	return hex.EncodeToString(saltBytes)
}

func GeneratePassword(salt, password string) string {
	df := pbkdf2.Key([]byte(password), []byte(salt), 10000, 512, sha512.New)
	cipherText := hex.EncodeToString(df)
	return cipherText
}

func NewVerifyPassword(password, storedpassword, salt string) bool {
	nPassword := GeneratePassword(salt, password)
	return nPassword == storedpassword
}
