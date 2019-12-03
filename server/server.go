package server

import (
	"encoding/json"
	"github.com/ipan97/go-dig-sample/config"
	"github.com/ipan97/go-dig-sample/service"
	"net/http"
)

type Server struct {
	config        *config.Config
	personService *service.PersonService
}

func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/people", s.handleGetAllPeople)
	return mux
}

func (s *Server) handleGetAllPeople(w http.ResponseWriter, r *http.Request) {
	people := s.personService.FindAll()
	bytes, _ := json.Marshal(people)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(bytes)
}

func (s *Server) Run() {
	httpServer := &http.Server{
		Addr:    ":" + s.config.Port,
		Handler: s.Handler(),
	}

	_ = httpServer.ListenAndServe()
}

func NewServer(config *config.Config, service *service.PersonService) *Server {
	return &Server{
		config:        config,
		personService: service,
	}
}
