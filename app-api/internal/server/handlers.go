package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"task-api/internal/model"
	"time"
)

const (
	serializeError = "serialize error: %s"
	dbError        = "database error: %s"
	dateFormat     = "2006-01-02 15:04:05 -0700"
)

func (s *Server) errorLog(writer http.ResponseWriter, error string, code int) {
	http.Error(writer, error, code)
	log.Printf(error)
}

func (s *Server) createTaskHandler(writer http.ResponseWriter, request *http.Request) {
	data := struct {
		Name        string `json:"Name"`
		Description string `json:"Description"`
		EndDate     string `json:"End_date"`
		UserName    string `json:"User_name"`
		StartDate   string `json:"Start_date"`
	}{}

	if err := json.NewDecoder(request.Body).Decode(&data); err != nil {
		s.errorLog(writer, fmt.Sprintf(serializeError, err), http.StatusBadRequest)
		return
	}
	log.Printf("incoming task: %v", data)

	newTask := model.NewTask()
	var err error

	newTask.UserId, err = s.dbHandler.GetUserIdByName(data.UserName)
	if err != nil {
		s.errorLog(writer, fmt.Sprintf(serializeError, err), http.StatusBadRequest)
		return
	}

	newTask.EndDate, err = time.Parse(dateFormat, data.EndDate)
	if err != nil {
		s.errorLog(writer, fmt.Sprintf("end date parse error: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	newTask.StartDate, err = time.Parse(dateFormat, data.StartDate)
	if err != nil {
		s.errorLog(writer, fmt.Sprintf("start date parce error: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	newTask.Name = data.Name
	newTask.Description = data.Description
	if err := s.dbHandler.CreateTask(*newTask); err != nil {
		s.errorLog(writer, fmt.Sprintf("database error: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func (s *Server) getTasksHandler(writer http.ResponseWriter, request *http.Request) {
	queryValues := request.URL.Query()

	name := queryValues.Get("name")
	log.Printf("query name: %s", name)
	if name == "" {
		s.errorLog(writer, fmt.Sprintf("empty name"), http.StatusBadRequest)
		return
	}

	userId, err := s.dbHandler.GetUserIdByName(name)
	if err != nil {
		s.errorLog(writer, fmt.Sprintf("db error: %s", err.Error()), http.StatusBadRequest)
		return
	}

	tasks, err := s.dbHandler.GetUserTasks(userId)
	if err != nil {
		s.errorLog(writer, fmt.Sprintf("db error: %s", err.Error()), http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(writer).Encode(tasks); err != nil {
		s.errorLog(writer, fmt.Sprintf("serialize error: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
}
