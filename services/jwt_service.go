package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type jwtService struct {
	secretKey string
	issure string
}

func NewJWTService(secretKey string, issure string) *jwtService {
	return &jwtService{
    secretKey: secretKey,
    issure: issure,
  }
}

type Clain struct {
	Sum uint `json:"sum"`
	jwt.StandardClaims
}

func (s *jwtService) GenerateToken(id uint) (string, error) {
	clain := &Clain{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer: s.issure,
			IssuedAt: time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clain)

	t, err := token.SignedString([]byte(s.secretKey))
	if err!= nil {
    return "", err
  }

	return t, nil
}

func(s *jwtService) ValidateToken (token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("Invalid token: %v", token)
		}
		return []byte(s.secretKey), nil
	})

	return err == nil
}