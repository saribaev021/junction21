package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type task struct {
	Name     string `json:"Name"`
	UserName string `json:"User_name"`
}

func (s *Server) deleteFinishedTask(writer http.ResponseWriter, request *http.Request) {
	var taskData task

	if err := json.NewDecoder(request.Body).Decode(&taskData); err != nil {
		s.errorLog(writer, fmt.Sprintf("serialize error: %s", err), http.StatusBadRequest)
		return
	}

	log.Printf("incoming data: %v", taskData)

	user, err := s.dbHandler.GetUserByName(taskData.UserName)
	if err != nil {
		s.errorLog(writer, fmt.Sprintf("database GetUserByName error: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	modelTask, err := s.dbHandler.GetTaskByUser(user.Id, taskData.Name)
	if err != nil {
		s.errorLog(writer, fmt.Sprintf("database GetTaskByUser error: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	diff := (time.Now().UnixNano()/int64(time.Hour) - modelTask.EndDate.UnixNano()/int64(time.Hour)) / 24
	if diff > 0 {
		modelTask.Xp -= int(diff * 100)
	}
	user.Xp += modelTask.Xp

	if modelTask.Xp > 0 {
		if err := s.dbHandler.UpdateUserXP(user.Id, user.Xp); err != nil {
			s.errorLog(writer, fmt.Sprintf("database UpdateUserXP error: %s", err.Error()), http.StatusInternalServerError)
			return
		}
	}

	if err := s.dbHandler.DeleteTaskByUser(modelTask.Name, user.Id); err != nil {
		s.errorLog(writer, fmt.Sprintf("database DeleteTaskByUser error: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	writer.Write([]byte("task deleted"))
}

func (s *Server) deleteTaskGiveUp(writer http.ResponseWriter, request *http.Request) {
	var taskData task

	if err := json.NewDecoder(request.Body).Decode(&taskData); err != nil {
		s.errorLog(writer, fmt.Sprintf("serialize error: %s", err), http.StatusBadRequest)
		return
	}

	log.Printf("incoming data: %v", taskData)
	userId, err := s.dbHandler.GetUserIdByName(taskData.UserName)
	if err != nil {
		s.errorLog(writer, fmt.Sprintf("database GetUserIdByName error: %s", err), http.StatusBadRequest)
		return
	}

	if err := s.dbHandler.DeleteTaskByUser(taskData.Name, userId); err != nil {
		s.errorLog(writer, fmt.Sprintf("database DeleteTaskByUser error: %s", err.Error()), http.StatusInternalServerError)
		return
	}

}
