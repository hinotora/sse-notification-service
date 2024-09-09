package connection

type Application struct {
	Id string `json:"id"`
}

func CreateApplication(id string) *Application {
	return &Application{
		Id: id,
	}
}

func (a Application) GetId() string {
	return a.Id
}
