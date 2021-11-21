package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func (s *Server) deleteFinishedTask(writer http.ResponseWriter, request *http.Request) {
	task := struct {
		Name     string `json:"Name"`
		UserName string `json:"User_name"`
	}{}
	if err := json.NewDecoder(request.Body).Decode(&task); err != nil {
		s.errorLog(writer, fmt.Sprintf("serialize error: %s", err), http.StatusBadRequest)
		return
	}

	user, err := s.dbHandler.GetUserByName(task.UserName)
	if err != nil {
		s.errorLog(writer, fmt.Sprintf("database error: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	modelTask, err := s.dbHandler.GetTaskByUser(user.Id, task.Name)
	if err != nil {
		s.errorLog(writer, fmt.Sprintf("database error: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	diff := (time.Now().UnixNano() / int64(time.Hour) - modelTask.EndDate.UnixNano() / int64(time.Hour)) / 24
	if diff > 0 {
		modelTask.Xp -= int(diff * 60)
	}
	user.Xp += modelTask.Xp

	if err := s.dbHandler.UpdateUserXP(user.Id, user.Xp); err != nil {
		s.errorLog(writer, fmt.Sprintf("database error: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	if err := s.dbHandler.DeleteTaskByUser(modelTask.Name, user.Id); err != nil {
		s.errorLog(writer, fmt.Sprintf("database error: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	log.Printf("incoming task: %v", task)
}
