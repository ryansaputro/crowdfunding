package auth

import (
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

type Service interface {
	GenerateToken(userID int) (string, error)
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	// payload bisa disebut payload atau claim
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	secretKeyEnv := os.Getenv("SECRET_KEY")
	var secretKey = []byte(secretKeyEnv)
	signedToken, err := token.SignedString(secretKey)

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
