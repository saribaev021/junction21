package model

import "time"

type Task struct {
	Id          int
	Name        string
	StartDate   time.Time
	EndDate     time.Time
	Description string
	UserId      int
	Xp          int
}

type User struct {
	Id          int
	Name        string    		`json:"Name"`
	Xp 			int             `json:"Xp"`
}

func NewUser() *User {
	return &User{}
}

func NewTask() *Task {
	return &Task{}
}
