package services

import (
	"errors"
	"pasour/internal/domain/commands"
	"pasour/internal/domain/dtos"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenService struct {
	Secret        []byte
	TokenDuration time.Duration
}

type claims struct {
	sub string
	jwt.RegisteredClaims
}

func NewTokenService(secret string, tokenDuration time.Duration) *TokenService {
	return &TokenService{
		Secret:        []byte(secret),
		TokenDuration: tokenDuration,
	}
}

func (ts *TokenService) Encode(cmd *commands.TokenEncodeCmd) (*dtos.TokenDTO, error) {
	now := time.Now()
	exp := now.Add(ts.TokenDuration)
	c := claims{
		sub: cmd.Sub,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	tokenStr := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	signedToken, err := tokenStr.SignedString(ts.Secret)
	if err != nil {
		return nil, err
	}

	token := dtos.NewTokenDTO(signedToken)
	return token, nil

}

func (ts *TokenService) Decode(token string) (string, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token method")
		}
		return ts.Secret, nil
	}

	parsedToken, err := jwt.ParseWithClaims(token, &claims{}, keyFunc)
	if err != nil || !parsedToken.Valid {
		return "", errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(*claims)
	if !ok {
		return "", errors.New("invalid token claims")
	}

	return claims.sub, nil
}

func (ts *TokenService) GetTokenFromHeader(authHeader string) (string, error) {
	if authHeader == "" {
		return "", errors.New("Authorization header is required")
	}
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 {
		return "", errors.New("Authorization header is invalid")
	}
	if parts[0] != "Bearer" {
		return "", errors.New("Authorization header is invalid")
	}

	return parts[1], nil
}
