package services

import (
	"pasour/internal/domain/commands"
	"pasour/internal/domain/errors"
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

func (ts *TokenService) Encode(cmd *commands.TokenEncodeCmd) (string, *errors.DomainErr) {
	now := time.Now()
	exp := now.Add(ts.TokenDuration)
	c := claims{
		sub: cmd.Sub,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	signedToken, err := token.SignedString(ts.Secret)
	if err != nil {
		return "", errors.NewInternalErr("error signing token")
	}

	return signedToken, nil

}

func (ts *TokenService) Decode(token string) (string, *errors.DomainErr) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.NewUnAuthorizedErr("unexpected signing method")
		}
		return ts.Secret, nil
	}

	parsedToken, err := jwt.ParseWithClaims(token, &claims{}, keyFunc)
	if err != nil || !parsedToken.Valid {
		return "", errors.NewUnAuthorizedErr("invalid token")
	}

	claims, ok := parsedToken.Claims.(*claims)
	if !ok {
		return "", errors.NewUnAuthorizedErr("invalid token")
	}

	return claims.sub, nil
}

func (ts *TokenService) GetTokenFromHeader(authHeader string) (string, *errors.DomainErr) {
	if authHeader == "" {
		return "", errors.NewUnAuthorizedErr("Authorization header is required")
	}
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 {
		return "", errors.NewUnAuthorizedErr("Authorization header is invalid")
	}
	if parts[0] != "Bearer" {
		return "", errors.NewUnAuthorizedErr("Authorization header is invalid")
	}

	return parts[1], nil
}
