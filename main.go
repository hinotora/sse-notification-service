package main

import (
	"fmt"

	"github.com/hinotora/sse-notification-service/internal"
	"github.com/hinotora/sse-notification-service/internal/config"
	"github.com/hinotora/sse-notification-service/internal/logger"
	"github.com/hinotora/sse-notification-service/internal/redis"
	"github.com/hinotora/sse-notification-service/internal/repository"
	"github.com/hinotora/sse-notification-service/internal/router"
)

var Logger *logger.Logger = nil

func main() {
	c, err := config.Load()

	if err != nil {
		panic(err)
	}

	Logger := logger.New()

	internal.PrintStart()
	internal.PrintConf(Logger, c)

	_, err = redis.Load(c)

	if err != nil {
		Logger.Fatal(fmt.Sprintf("Error connecting to Redis: %s", err))
	}

	repository.Init()

	err = router.Run()

	Logger.Fatal(fmt.Sprintf("Server error: %s", err))
}
