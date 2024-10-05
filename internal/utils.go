package internal

import (
	"fmt"
	"os"

	"github.com/hinotora/sse-notification-service/internal/config"
	"github.com/hinotora/sse-notification-service/internal/logger"
)

func PrintStart() {
	ver := "unknown"

	if buf, err := os.ReadFile("version"); err != nil {
		ver = "err_version_not_found"
	} else {
		ver = string(buf)
	}

	fmt.Println()
	fmt.Println("Go Websocket notification microservice")
	fmt.Println("Visit: https://github.com/hinotora/sse-notification-service")
	fmt.Printf("Version: %s \n", ver)
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
