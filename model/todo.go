package model

type Todo struct {
	Id     int    `json:"id"`
	UserId string `json:"userId"`
	Todo   string `json:"todo"`
	IsDone bool   `json:"isDone"`
}
