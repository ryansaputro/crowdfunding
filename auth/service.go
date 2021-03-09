package auth

import (
	"errors"
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) secretToken() ([]byte, error) {
	secretKeyEnv := os.Getenv("SECRET_KEY")
	var secretKey = []byte(secretKeyEnv)

	return secretKey, nil
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	// payload bisa disebut payload atau claim
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	secretKey, err := s.secretToken()
	if err != nil {
		return "error", errors.New("token tidak valid")
	}

	secretKey = []byte(secretKey)

	signedToken, err := token.SignedString(secretKey)

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {

	secretKey, err := s.secretToken()
	secretKey = []byte(secretKey)

	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			fmt.Println("token tidak valid")
			return nil, errors.New("token tidak valid")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		fmt.Println("token tidak valid: " + err.Error())
		return token, err
	}

	return token, nil
}
