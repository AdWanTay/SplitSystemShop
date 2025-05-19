package utils

import (
	"SplitSystemShop/internal/config"
	"SplitSystemShop/internal/models"
	"encoding/base64"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"gopkg.in/gomail.v2"
	"os"
	"path/filepath"
	"regexp"
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

// Сохраняет base64 в файл, возвращает URL
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

// Заменяет все base64-картинки в HTML на URL
func ReplaceBase64ImagesInHTML(html string) (string, error) {
	re := regexp.MustCompile(`(?i)<img[^>]+src="(data:image/[^;]+;base64,[^"]+)"`)
	matches := re.FindAllStringSubmatch(html, -1)

	updatedHTML := html

	for _, match := range matches {
		base64Str := match[1]

		url, err := SaveBase64Image(base64Str)
		if err != nil {
			return "", err
		}

		updatedHTML = strings.ReplaceAll(updatedHTML, base64Str, url)
	}

	return updatedHTML, nil
}

func SendNewOrderNotification(to string, order *models.Order, cfg *config.Config) error {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", cfg.SMTP.User, cfg.SMTP.From)
	m.SetHeader("To", to)
	m.SetHeader("Subject", fmt.Sprintf("🛒 Новый заказ №%d", order.ID))

	body := fmt.Sprintf(`
		<h2 style="color:#333;">Вы получили новый заказ №%d</h2>
		<p><strong>Клиент:</strong> %s</p>
		<p><strong>Дата заказа:</strong> %s</p>
		<p><strong>Общая сумма:</strong> %.2f ₽</p>
		<p><strong>Статус заказа:</strong> %s</p>

		<h3 style="margin-top:20px;">Состав заказа:</h3>
		<ul style="padding-left:20px;">
	`, order.ID, order.User.Email, order.CreatedAt.Format("02.01.2006 15:04"), float64(order.TotalPrice)/100, order.Status)

	for _, item := range order.SplitSystems {
		body += fmt.Sprintf("<li><strong>%s</strong> — %.2f ₽</li>", item.Title, float64(item.Price)/100)
	}

	body += `</ul>`

	body += `<p style="margin-top:20px; font-style: italic; color: #777;">Это письмо сгенерировано автоматически. Пожалуйста, не отвечайте на него.</p>`

	m.SetBody("text/html", body)

	d := gomail.NewDialer(cfg.SMTP.Host, cfg.SMTP.Port, cfg.SMTP.User, cfg.SMTP.Password)
	return d.DialAndSend(m)
}

func SendOrderStatusUpdateNotification(to string, order *models.Order, cfg *config.Config) error {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", cfg.SMTP.User, cfg.SMTP.From)
	m.SetHeader("To", to)
	m.SetHeader("Subject", fmt.Sprintf("🔄 Обновление статуса заказа №%d", order.ID))

	body := fmt.Sprintf(`
		<h2 style="color:#333;">Статус вашего заказа №%d обновлён</h2>
		<p><strong>Новый статус:</strong> %s</p>
		<p><strong>Дата изменения:</strong> %s</p>
		<p><strong>Сумма заказа:</strong> %.2f ₽</p>

	`, order.ID, order.Status, time.Now().Format("02.01.2006 15:04"), float64(order.TotalPrice)/100)

	body += `<p style="margin-top:20px; font-style: italic; color: #777;">Если вы не делали этот заказ — пожалуйста, свяжитесь с нашей службой поддержки.</p>`
	m.SetBody("text/html", body)
	d := gomail.NewDialer(cfg.SMTP.Host, cfg.SMTP.Port, cfg.SMTP.User, cfg.SMTP.Password)
	return d.DialAndSend(m)
}
