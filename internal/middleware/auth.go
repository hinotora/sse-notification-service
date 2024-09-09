package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/hinotora/sse-notification-service/internal/config"

	"github.com/golang-jwt/jwt/v5"
)

func ExtractToken(request *http.Request) string {
	headerVal := request.Header.Get("Authorization")

	if len(headerVal) != 0 {
		splitedVal := strings.Split(headerVal, " ")

		if splitedVal[0] == "Bearer" && len(splitedVal[1]) != 0 {
			return splitedVal[1]
		}
	}

	queryVal := request.URL.Query().Get("Authorization")

	if len(queryVal) != 0 {
		return queryVal
	}

	return ""
}

func ValidateToken(config *config.Config, tokenstring string) (jwt.MapClaims, error) {
	if len(tokenstring) == 0 {
		return nil, errors.New("empty token")
	}

	// Парсим токен
	token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {

		// Нужно проверить, что метод подписи - HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(config.JWT.SecretKey), nil
	})

	// Проверяем ошибки валидации токена, если есть, рвем соединение
	if err != nil {
		return nil, fmt.Errorf("validation: %s", err)
	}

	// Пытаемся распарсить данные из токена
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, fmt.Errorf("claims validation: %s", err)
	}

	_, err = claims.GetIssuer()

	if err != nil {
		return nil, errors.New("issuer not found")
	}

	_, err = claims.GetSubject()

	if err != nil {
		return nil, errors.New("subject not found")
	}

	return claims, nil
}
