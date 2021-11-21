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
	dateFormat = "2006-01-02 15:04:05 -0700"
)

func (s *Server) errorLog(writer http.ResponseWriter, error string, code int) {
	http.Error(writer, error, code)
	log.Printf(error)
}

type getTasksEncoder struct {
	Name        string `json:"Name"`
	StartDate   string `json:"Start_date"`
	EndDate     string `json:"End_date"`
	Description string `json:"Description"`
}

type createTaskDecoder struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
	EndDate     string `json:"End_date"`
	UserName    string `json:"User_name"`
	StartDate   string `json:"Start_date"`
}

func (s *Server) xpCalculator(startDate time.Time, endDate time.Time) int {
	return int(endDate.Sub(startDate).Hours()/24) * 100
}

func (s *Server) createTaskHandler(writer http.ResponseWriter, request *http.Request) {
	data := createTaskDecoder{}

	if err := json.NewDecoder(request.Body).Decode(&data); err != nil {
		s.errorLog(writer, fmt.Sprintf("serialize error: %s", err), http.StatusBadRequest)
		return
	}
	log.Printf("incoming task: %v", data)

	newTask := model.NewTask()
	var err error

	newTask.UserId, err = s.dbHandler.GetUserIdByName(data.UserName)
	if err != nil {
		s.errorLog(writer, fmt.Sprintf("serialize error: %s", err), http.StatusBadRequest)
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
	newTask.Xp = s.xpCalculator(newTask.StartDate, newTask.EndDate)
	if err := s.dbHandler.CreateTask(*newTask); err != nil {
		s.errorLog(writer, fmt.Sprintf("database error: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	writer.Write([]byte("task created"))
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
		s.errorLog(writer, fmt.Sprintf("db GetUserTasks error: %s", err.Error()), http.StatusBadRequest)
		return
	}

	jsons := make([]getTasksEncoder, len(tasks))

	for i, task := range tasks {
		jsons[i] = getTasksEncoder{
			Name:        task.Name,
			EndDate:     task.EndDate.Format(dateFormat),
			StartDate:   task.StartDate.Format(dateFormat),
			Description: task.Description,
		}
	}

	if err := json.NewEncoder(writer).Encode(jsons); err != nil {
		s.errorLog(writer, fmt.Sprintf("serialize error: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
}
