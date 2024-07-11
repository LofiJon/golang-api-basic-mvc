package dto

type TaskDto struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}
