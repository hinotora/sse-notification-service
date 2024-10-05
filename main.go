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

func main() {
	c, err := config.Load()

	if err != nil {
		panic(err)
	}

	logger.Instance = logger.New()

	internal.PrintStart()
	internal.PrintConf(logger.Instance, c)

	_, err = redis.Load(c)

	if err != nil {
		logger.Instance.Fatal(fmt.Sprintf("Error connecting to Redis: %s", err))
	}

	repository.Init()

	fmt.Print("\n\n")
	logger.Instance.Info("Init sequence completed. Ready to accept connections. \n\n")

	err = router.Run()

	logger.Instance.Fatal(fmt.Sprintf("Server error: %s", err))
}
