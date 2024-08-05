package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const minSecretKeySize = 32

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d character", minSecretKeySize)
	}

	return &JWTMaker{secretKey}, nil
}

func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", payload, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, MyCustomClaims{
		Payload:          *payload,
		RegisteredClaims: jwt.RegisteredClaims{},
	})
	token, err :=  jwtToken.SignedString([]byte(maker.secretKey))
	return token, payload, err
}

func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	Keyfunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(maker.secretKey), nil
	}
	jwtToken, err := jwt.ParseWithClaims(token, &MyCustomClaims{}, Keyfunc)
	if err != nil {
		return nil, ErrInvalidToken
	}
	if claims, ok := jwtToken.Claims.(*MyCustomClaims); ok && jwtToken.Valid {
		if claims.Payload.ExpiredAt.After(time.Now()) {
			return &claims.Payload, nil
		}
		return nil, ErrExpiredToken

	}

	return nil, ErrInvalidToken

}
