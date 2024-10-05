package router

import (
	"net/http"

	"github.com/hinotora/sse-notification-service/internal/controller"
	"github.com/hinotora/sse-notification-service/internal/middleware"
)

func Run() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.Header().Add("Location", "/health")
		w.WriteHeader(http.StatusMovedPermanently)

	})

	// Хэндлер эндпоинта проверки жив ли сервис
	mux.HandleFunc("GET /health", controller.Health)

	// Хэндлер непосредственно канала
	mux.Handle("/sse", middleware.CheckAuth(middleware.Repository(http.HandlerFunc(controller.OpenSSE))))

	// Хэндлеры работы с клиентами
	mux.Handle("GET /connections/{application_id}", middleware.CheckAuth(http.HandlerFunc(controller.GetApplicationConnections)))
	mux.Handle("GET /connections/{application_id}/{user_id}", middleware.CheckAuth(http.HandlerFunc(controller.GetUserConnections)))

	err := http.ListenAndServe(":80", mux)

	return err
}
