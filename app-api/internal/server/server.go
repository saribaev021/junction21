package server

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"task-api/internal/model"
)

type handler interface {
	CreateTask(model.Task) error
	GetUserIdByName(string) (int, error)
	GetUserByName(string) (model.User, error)
	GetUserTasks(int) ([]model.Task, error)
	GetTaskByUser(userId int, taskName string) (model.Task, error)
	UpdateUserXP(userId int, xp int) error
	DeleteTaskByUser(name string, userId int) error
}

type Server struct {
	router    *chi.Mux
	dbHandler handler
}

func NewServer(h handler) *Server {
	return &Server{
		router:    chi.NewRouter(),
		dbHandler: h,
	}
}

func (s *Server) InitSever() {
	log.Println("initing server...")

	s.router.Post("/create/task", s.createTaskHandler)
	s.router.Get("/get/tasks", s.getTasksHandler)
	s.router.Post("/delete/task/finished", s.deleteFinishedTask)
	s.router.Post("/delete/task/giveup", s.deleteTaskGiveUp)

	s.router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("ping"))
	})
}

func (s *Server) Listen(addr string) error {
	log.Printf("listening on address %s", addr)

	return http.ListenAndServe(addr, s.router)
}
