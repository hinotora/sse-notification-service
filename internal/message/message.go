package message

import (
	"github.com/google/uuid"
)

type Message struct {
	Id    string            `json:"id"`
	Mtype string            `json:"type"`
	Data  map[string]string `json:"data"`
}

var (
	TYPEPING  = "ping"
	TYPEEVENT = "event"
)

func NewMessage(data map[string]string, mtype string) *Message {
	return &Message{
		Id:    uuid.New().String(),
		Data:  data,
		Mtype: mtype,
	}
}
