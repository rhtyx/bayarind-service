package token

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	Jwt *JWT
)

func InitJWT() {
	privateKey, err := os.ReadFile("./cert/id_rsa.pri")
	if err != nil {
		log.Fatal(err)
	}

	publicKey, err := os.ReadFile("./cert/id_rsa.pub")
	if err != nil {
		log.Fatal(err)
	}

	Jwt = &JWT{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}
}

type JWTService interface {
	CreateToken(userID int64, createdAt time.Time, duration time.Duration) (string, error)
	ValidateToken(token string) (*Claims, error)
}

type JWT struct {
	PrivateKey []byte
	PublicKey  []byte
}

func (j JWT) CreateToken(userID int64, createdAt time.Time, duration time.Duration) (string, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM(j.PrivateKey)
	if err != nil {
		return "", err
	}

	claims, err := NewClaims(userID, createdAt, duration)
	if err != nil {
		return "", err
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (j JWT) ValidateToken(token string) (*Claims, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM(j.PublicKey)
	if err != nil {
		return nil, err
	}

	parsedToken, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return key, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(*Claims)
	if !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	if claims.ExpiresAt.Before(time.Now()) {
		return nil, fmt.Errorf("token expired")
	}

	return claims, nil
}
