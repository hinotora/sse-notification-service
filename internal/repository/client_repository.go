package repository

import (
	"context"
	"errors"
	"sync"

	"github.com/hinotora/sse-notification-service/internal/connection"
	"github.com/hinotora/sse-notification-service/internal/redis"
)

type ClientRepository struct {
	Connections map[string]map[string]map[string]*connection.Connection

	clientMutex sync.Mutex
}

var repo *ClientRepository = nil

func Init() *ClientRepository {
	repo = &ClientRepository{}

	repo.Connections = make(map[string]map[string]map[string]*connection.Connection)

	return repo
}

func AddClient(c *connection.Connection) (int, error) {
	repo.clientMutex.Lock()
	defer repo.clientMutex.Unlock()

	if repo.Connections[c.Application.GetId()] == nil {
		repo.Connections[c.Application.Id] = make(map[string]map[string]*connection.Connection)
	}

	if repo.Connections[c.Application.GetId()][c.User.GetId()] == nil {
		repo.Connections[c.Application.GetId()][c.User.GetId()] = make(map[string]*connection.Connection)
	}

	repo.Connections[c.Application.GetId()][c.User.GetId()][c.GetId()] = c

	err := redis.GetInstance().HMSet(context.TODO(), c.GetHashName(), map[string]interface{}{
		"application_id": c.Application.GetId(),
		"user_id":        c.User.GetId(),
		"channel_id":     c.GetChannelName(),
	}).Err()

	if err != nil {
		return 0, errors.New("redis: error while creating user hash")
	}

	return len(repo.Connections[c.Application.GetId()][c.User.GetId()]), nil
}

func DelClient(c *connection.Connection) {
	repo.clientMutex.Lock()
	defer repo.clientMutex.Unlock()

	delete(repo.Connections[c.Application.GetId()][c.User.GetId()], c.GetId())

	if len(repo.Connections[c.Application.GetId()][c.User.GetId()]) == 0 {
		delete(repo.Connections[c.Application.GetId()], c.User.GetId())

		redis.GetInstance().Del(context.TODO(), c.GetHashName())

		if len(repo.Connections[c.Application.GetId()]) == 0 {
			delete(repo.Connections, c.Application.GetId())
		}
	}
}

func GetApplicationConnections(applicationId string) map[string]map[string]*connection.Connection {

	if repo.Connections[applicationId] != nil {
		return repo.Connections[applicationId]
	}

	return nil
}

func GetUsersConnections(applicationId string, userId string) map[string]*connection.Connection {

	if repo.Connections[applicationId][userId] != nil {
		return repo.Connections[applicationId][userId]
	}

	return nil
}
