package connection

import (
	"context"
	"fmt"
	"time"

	"github.com/hinotora/sse-notification-service/internal/message"

	"github.com/google/uuid"
)

type Connection struct {
	Id string `json:"id"`

	CreatedAt  int64  `json:"created_at"`
	RemoteAddr string `json:"remote_addr"`

	Application *Application `json:"application"`
	User        *User        `json:"user"`

	PingCh      chan bool            `json:"-"`
	BroadcastCh chan message.Message `json:"-"`
	CloseCh     <-chan struct{}      `json:"-"`
}

func Create(app *Application, user *User, ctx context.Context) *Connection {
	c := &Connection{}

	c.Id = uuid.New().String()
	c.CreatedAt = time.Now().Unix()

	c.PingCh = make(chan bool)
	c.BroadcastCh = make(chan message.Message)
	c.CloseCh = ctx.Done()

	c.Application = app
	c.User = user

	return c
}

func (c *Connection) GetId() string {
	return c.Id
}

func (c *Connection) GetChannelName() string {
	return fmt.Sprintf("%s:channel:%s", c.Application.GetId(), c.User.GetId())
}

func (c *Connection) GetHashName() string {
	return fmt.Sprintf("%s:client:%s", c.Application.GetId(), c.User.GetId())
}
