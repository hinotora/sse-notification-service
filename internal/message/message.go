package message

type Message struct {
	Id    string            `json:"id"`
	Mtype string            `json:"type"`
	Data  map[string]string `json:"data"`
}

func NewMessage(id string, data map[string]string, mtype string) *Message {
	return &Message{
		Id:    id,
		Data:  data,
		Mtype: mtype,
	}
}
