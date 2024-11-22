package models

type Task struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Finished bool   `json:"finished"`
}
