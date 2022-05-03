package Entities

type Entity struct {
	Id int `json:"Id"`
}

func (this Entity) getId() int {
	return this.Id
}
