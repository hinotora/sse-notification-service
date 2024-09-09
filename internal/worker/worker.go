package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hinotora/sse-notification-service/internal/connection"
	"github.com/hinotora/sse-notification-service/internal/logger"
	"github.com/hinotora/sse-notification-service/internal/message"
	"github.com/hinotora/sse-notification-service/internal/redis"
)

func Writer(conn *connection.Connection, ctx context.Context) {
	logger := ctx.Value("logger").(*logger.Logger)

	channelName := conn.GetChannelName()

	pubsub := redis.GetInstance().Subscribe(context.TODO(), channelName)

	defer pubsub.Close()

	logger.Info(fmt.Sprintf("Listening channel: %s", channelName))

	logger.Debug("Writer loop started")

writer:
	for {
		select {
		case rawMsg := <-pubsub.Channel():
			msg := &message.Message{}

			err := json.Unmarshal([]byte(rawMsg.Payload), msg)

			if err != nil {
				logger.Debug(err)
				continue
			}

			conn.BroadcastCh <- *msg

		case <-conn.PingCh:
			conn.BroadcastCh <- *message.NewMessage(nil, message.TYPEPING)
		case <-conn.CloseCh:
			break writer

		}

	}

	logger.Debug("Writer loop stopped")

}

func Pinger(conn *connection.Connection, ctx context.Context) {
	logger := ctx.Value("logger").(*logger.Logger)

	logger.Debug("Pinger loop started")

	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

loop:
	for {
		select {
		case <-conn.CloseCh:
			break loop
		case <-ticker.C:
			conn.PingCh <- true
		}
	}

	logger.Debug("Pinger loop stopped")
}
