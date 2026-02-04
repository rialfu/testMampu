package service

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService interface {
	GenerateAccessToken(userId uint64) string
	ValidateToken(token string) (*jwt.Token, error)
	GetUserIDByToken(token string) (string, error)
}

type jwtCustomClaim struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

type jwtService struct {
	secretKey     string
	issuer        string
	accessExpiry  time.Duration
	refreshExpiry time.Duration
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey:     getSecretKey(),
		issuer:        "Template",
		accessExpiry:  time.Hour * 24,
		refreshExpiry: time.Hour * 24 * 7,
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		secretKey = "Template"
	}
	return secretKey
}

func (j *jwtService) GenerateAccessToken(userId uint64) string {
	claims := jwtCustomClaim{
		fmt.Sprintf("%d", userId),
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.accessExpiry)),
			Issuer:    j.issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tx, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		log.Println(err)
	}
	return tx
}

func (j *jwtService) parseToken(t_ *jwt.Token) (any, error) {
	if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method %v", t_.Header["alg"])
	}
	return []byte(j.secretKey), nil
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, j.parseToken)
}

func (j *jwtService) GetUserIDByToken(token string) (string, error) {
	tToken, err := j.ValidateToken(token)
	if err != nil {
		return "", err
	}

	claims := tToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id, nil
}
