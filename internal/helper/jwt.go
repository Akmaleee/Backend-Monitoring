package helper

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"it-backend/internal/model/entity"
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type ClaimToken struct {
	UserID   int    `json:"user_id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var MapTypeToken = map[string]time.Duration{
	"access_token":  time.Hour * 3,
	"refresh_token": time.Hour * 72,
}

func GenerateToken(ctx context.Context, user entity.User, tokenType string, now time.Time) (string, error) {
	config, err := GetConfig()
	if err != nil {
		log.Panic(err)
	}

	// claimToken := ClaimToken{
	// 	UserID:   user.ID,
	// 	Name:     user.Name,
	// 	Username: user.Username,
	// 	RegisteredClaims: jwt.RegisteredClaims{
	// 		Issuer:    config.APP_NAME,
	// 		IssuedAt:  jwt.NewNumericDate(now),
	// 		ExpiresAt: jwt.NewNumericDate(now.Add(MapTypeToken[tokenType])),
	// 	},
	// }

	// New Claim Token
	claimToken := jwt.MapClaims{
		"user": map[string]interface{}{
			"id":       user.ID,
			"name":     user.Name,
			"username": user.Username,
			"role":     user.RoleUser,
		},
		"exp": now.UTC().Add(time.Duration(config.JWT_EXPIRE) * time.Second).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimToken)

	resultToken, err := token.SignedString([]byte(config.JWT_SECRET))
	if err != nil {
		return resultToken, fmt.Errorf("failed to generate token: %v", err)
	}
	return resultToken, nil
}

func ValidateToken(ctx context.Context, tokenStr string) (*ClaimToken, error) {
	config, err := GetConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %v", err)
	}

	jwtToken, err := jwt.ParseWithClaims(tokenStr, &ClaimToken{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(config.JWT_SECRET), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse jwt: %v", err)
	}

	claimToken, ok := jwtToken.Claims.(*ClaimToken)
	if !ok || !jwtToken.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claimToken, nil
}

func GetJWTPayload(ctx *fiber.Ctx) (map[string]interface{}, error) {
	authHeader := ctx.Get("Authorization")
	if authHeader == "" {
		return nil, fmt.Errorf("authorization header missing")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return nil, fmt.Errorf("invalid authorization header format")
	}

	token := parts[1]
	jwtParts := strings.Split(token, ".")
	if len(jwtParts) < 2 {
		return nil, fmt.Errorf("invalid JWT token")
	}

	payloadSegment := jwtParts[1]
	payloadBytes, err := base64.RawURLEncoding.DecodeString(payloadSegment)
	if err != nil {
		return nil, fmt.Errorf("failed to decode JWT payload: %w", err)
	}

	var payload map[string]interface{}
	if err := json.Unmarshal(payloadBytes, &payload); err != nil {
		return nil, fmt.Errorf("failed to parse JWT payload: %w", err)
	}

	return payload, nil
}

func GetUsernameFromPayload(ctx *fiber.Ctx) (string, error) {
	payload, err := GetJWTPayload(ctx)
	if err != nil {
		return "", err
	}

	userRaw, ok := payload["user"]
	if !ok {
		return "", errors.New("missing user field")
	}

	userMap, ok := userRaw.(map[string]interface{})
	if !ok {
		return "", errors.New("invalid user data")
	}

	username, ok := userMap["username"].(string)
	if !ok || username == "" {
		return "", errors.New("missing username")
	}

	return username, nil
}
