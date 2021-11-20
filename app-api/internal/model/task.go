package model

import "time"

type Task struct {
	Id          int
	Name        string
	StartDate   time.Time
	EndDate     time.Time
	Description string
	UserId      int
}

func NewTask() *Task {
	return &Task{}
}
