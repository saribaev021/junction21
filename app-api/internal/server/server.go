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
	GetUserTasks(int) ([]model.Task, error)
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
	s.router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Print("ruslan")
	})
}

func (s *Server) Listen(addr string) error {
	log.Printf("listening on address %s", addr)

	return http.ListenAndServe(addr, s.router)
}
