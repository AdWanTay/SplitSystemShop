package utils

import (
	"SplitSystemShop/internal/config"
	"SplitSystemShop/internal/models"
	"encoding/base64"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func GenerateJWT(user *models.User, cfg *config.Config) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(cfg.JWTSecret))
	return tokenString, nil
}

func ParseAndValidateJWT(tokenString string, cfg *config.Config) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(cfg.JWTSecret), nil
	})

	if err != nil || !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return 0, fmt.Errorf("invalid token claims")

	}
	userID, ok := claims["user_id"].(float64)

	return uint(userID), err
}

func ParseInt(v string) int {
	i, _ := strconv.Atoi(v)
	return i
}

func ParseUint(v string) uint {
	i, _ := strconv.ParseUint(v, 10, 64)
	return uint(i)
}

func ParseFloat(v string) float64 {
	f, _ := strconv.ParseFloat(v, 64)
	return f
}

func SaveBase64Image(data string) (string, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return "", fmt.Errorf("невалидный base64")
	}

	decoded, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return "", err
	}

	filename := fmt.Sprintf("%d.png", time.Now().UnixNano())
	path := filepath.Join("web", "static", "uploads", "article_images", filename)

	err = os.WriteFile(path, decoded, 0644)
	if err != nil {
		return "", err
	}

	return "/web/static/uploads/article_images/" + filename, nil
}
