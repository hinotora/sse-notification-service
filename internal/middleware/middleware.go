package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hinotora/sse-notification-service/internal/config"
	"github.com/hinotora/sse-notification-service/internal/connection"
	"github.com/hinotora/sse-notification-service/internal/logger"
	"github.com/hinotora/sse-notification-service/internal/repository"
)

func CheckAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Берем токен реквеста
		tokenString := ExtractToken(r)

		// Валидируем токен
		claims, err := ValidateToken(config.GetInstance(), tokenString)

		if err != nil {
			logger.Instance.Debug(fmt.Sprintf("token: %s", err))

			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		app := connection.CreateApplication(claims["iss"].(string))
		user := connection.CreateUser(claims["sub"].(string))

		conn := connection.Create(app, user, r.Context())

		logger := logger.New()
		logger.SetPrefix(fmt.Sprintf("[app:%s][usr:%s][cn:%s]", conn.Application.GetId(), conn.User.GetId(), conn.GetId()))

		ctx := r.Context()
		ctx = context.WithValue(ctx, "connection", conn)
		ctx = context.WithValue(ctx, "logger", logger)

		ctx, cancel := context.WithCancel(ctx)

		defer cancel()

		r = r.WithContext(ctx)

		// Продолжаем обработку соединения
		next.ServeHTTP(w, r)

		conn = nil
		app = nil
		user = nil
		claims = nil
		logger = nil
	})
}

func Repository(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := r.Context().Value("logger").(*logger.Logger)
		conn := r.Context().Value("connection").(*connection.Connection)

		logger.Debug("New authorized connection")

		count, err := repository.AddClient(conn)

		if err != nil {
			logger.Error(fmt.Sprintf("repository add error: %s", err))

			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		logger.Info(fmt.Sprintf("Client created. Same app/user connections: %d", count))

		defer repository.DelClient(conn)

		// Продолжаем обработку соединения
		next.ServeHTTP(w, r)

	})
}
