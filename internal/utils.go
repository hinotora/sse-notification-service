package internal

import (
	"fmt"

	"github.com/hinotora/sse-notification-service/internal/config"
	"github.com/hinotora/sse-notification-service/internal/logger"
)

func PrintStart() {
	fmt.Println()
	fmt.Println("Go Websocket notification microservice")
	fmt.Println("Visit: https://github.com/hinotora")
	fmt.Println("Version: 0.1.0")
}

func PrintConf(logger *logger.Logger, config *config.Config) {
	fmt.Println()

	logger.Info(fmt.Sprintf("Service running in <%s> mode", config.App.Mode))

	logger.Debug(fmt.Sprintf("APP_NAME = %s", config.App.Name))
	logger.Debug(fmt.Sprintf("APP_PORT = %s", config.App.Port))
	logger.Debug(fmt.Sprintf("REDIS_HOST = %s", config.Redis.Host))
	logger.Debug(fmt.Sprintf("REDIS_PORT = %s", config.Redis.Port))
	logger.Debug(fmt.Sprintf("JWT_SECRET = %s", config.JWT.SecretKey))

}
