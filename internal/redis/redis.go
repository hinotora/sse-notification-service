package redis

import (
	"context"
	"fmt"

	"github.com/hinotora/sse-notification-service/internal/config"

	rds "github.com/redis/go-redis/v9"
)

var instance *rds.Client

func GetInstance() *rds.Client {
	return instance
}

func Load(config *config.Config) (*rds.Client, error) {
	if instance != nil {
		return instance, nil
	}

	instance = rds.NewClient(&rds.Options{
		Addr: fmt.Sprintf("%s:%s", config.Redis.Host, config.Redis.Port),
	})

	if err := instance.Ping(context.TODO()).Err(); err != nil {
		return instance, err
	}

	return instance, nil
}
