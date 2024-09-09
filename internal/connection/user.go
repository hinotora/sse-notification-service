package connection

type User struct {
	Id string `json:"id"`
}

func CreateUser(id string) *User {
	return &User{
		Id: id,
	}
}

func (a User) GetId() string {
	return a.Id
}
