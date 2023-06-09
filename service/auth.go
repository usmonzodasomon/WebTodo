package service

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"time"
	"webtodo/pkg/repository"

	"github.com/dgrijalva/jwt-go"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId uint `json:"id"`
}

const signingKey = "kajsdljaskdja332$#"

func generatePasswordHash(pass string) string {
	hash := sha256.Sum256([]byte(pass))
	return fmt.Sprintf("%x", hash)
}

func GenerateToken(username, password string) (string, error) {
	id, err := repository.GetUserByUserAndPassword(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		id,
	})
	return token.SignedString([]byte(signingKey))
}

func ParseToken(tokenString string) (uint, error) {
	token, err := jwt.ParseWithClaims(tokenString, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type of *tokenClaims")
	}

	// if time.Now().Unix() > claims.ExpiresAt {
	// 	return 0, errors.New("token expired")
	// }

	return claims.UserId, nil
}
