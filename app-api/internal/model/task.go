package model

import "time"

type Task struct {
	Id          int
	Name        string    `json:"Name"`
	StartDate   time.Time `json:"Start_date"`
	EndDate     time.Time `json:"End_date"`
	Description string    `json:"Description"`
	UserId      int
}

func NewTask() *Task {
	return &Task{}
}
