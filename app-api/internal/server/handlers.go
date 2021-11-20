package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"task-api/internal/model"
	"time"
)

func (s *Server) createTaskHandler(writer http.ResponseWriter, request *http.Request) {
	data := struct {
		Name        string    `json:"Name"`
		Description string    `json:"Description"`
		EndDate     time.Time `json:"End_date"`
		UserName    string    `json:"User_name"`
		StartDate   time.Time `json:"Start_date"`
	}{}

	if err := json.NewDecoder(request.Body).Decode(&data); err != nil {
		http.Error(writer, fmt.Sprintf("serialize error: %s", err), http.StatusBadRequest)
		return
	}
	log.Printf("incoming task: %v", data)

	newTask := model.NewTask()
	var err error

	newTask.UserId, err = s.dbHandler.GetUserIdByName(data.UserName)
	if err != nil {
		http.Error(writer, fmt.Sprintf("database error: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	newTask.Name = data.Name
	newTask.EndDate = data.EndDate
	newTask.Description = data.Description
	newTask.StartDate = data.StartDate
	if err := s.dbHandler.CreateTask(*newTask); err != nil {
		http.Error(writer, fmt.Sprintf("database error: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func (s *Server) getTasksHandler(writer http.ResponseWriter, request *http.Request) {
	queryValues := request.URL.Query()

	name := queryValues.Get("name")
	log.Printf("query name: %s", name)
	if name == "" {
		http.Error(writer, fmt.Sprintf("empty name"), http.StatusBadRequest)
		return
	}

	userId, err := s.dbHandler.GetUserIdByName(name)
	if err != nil {
		http.Error(writer, fmt.Sprintf("db error: %s", err.Error()), http.StatusBadRequest)
		return
	}

	tasks, err := s.dbHandler.GetUserTasks(userId)
	if err != nil {
		http.Error(writer, fmt.Sprintf("db error: %s", err.Error()), http.StatusBadRequest)
		return
	}

	log.Println(tasks)
	if err := json.NewEncoder(writer).Encode(tasks); err != nil {
		http.Error(writer, fmt.Sprintf("serialize error: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
}
