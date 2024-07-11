package models

type Task struct {
	ID   uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}
