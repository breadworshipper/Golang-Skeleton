package jwthandler

import (
	"mm-pddikti-cms/internal/infrastructure/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

func GenerateTokenPairString(payload CostumClaimsPayload) (string, string, error) {
	accessClaims := CustomClaims{
		UserId: payload.UserId,
		Role:   payload.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   "user",
			Issuer:    config.Envs.App.Name,
			ExpiresAt: jwt.NewNumericDate(payload.AccessTokenExpiration),
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			NotBefore: jwt.NewNumericDate(time.Now().UTC()),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &accessClaims)
	accessTokenString, err := accessToken.SignedString([]byte(config.Envs.Guard.JwtPrivateKey))
	if err != nil {
		log.Error().Err(err).Msg("jwthandler::GenerateTokenString - Error while signing access token")
		return "", "", err
	}

	refreshClaims := CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   "user",
			Issuer:    config.Envs.App.Name,
			ExpiresAt: jwt.NewNumericDate(payload.RefreshTokenExpiration),
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			NotBefore: jwt.NewNumericDate(time.Now().UTC()),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &refreshClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(config.Envs.Guard.JwtPrivateKey))
	if err != nil {
		log.Error().Err(err).Msg("jwthandler::GenerateTokenString - Error while signing refresh token")
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

func ParseTokenString(tokenString string) (*CustomClaims, error) {
	claims := &CustomClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Envs.Guard.JwtPrivateKey), nil
	})
	if err != nil {
		log.Error().Err(err).Msg("jwthandler::ParseTokenString - Error while parsing token")
		return nil, err
	}

	if !token.Valid {
		log.Error().Msg("jwthandler::ParseTokenString - Invalid token")
		return nil, err
	}

	return claims, nil
}
