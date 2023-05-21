package httpserver

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog"
	"github.com/rs/zerolog/log"
	"github.com/taskBackend/config"
	"github.com/taskBackend/repository"
	"net/http"
)

type Server struct {
	cfg *config.Config
	db  *repository.PgRepository
}

func NewServer(
	cfg *config.Config,
	db *repository.PgRepository,
) *Server {
	return &Server{
		cfg: cfg,
		db:  db,
	}
}

func (s *Server) GetHandler() http.Handler {
	router := chi.NewRouter()
	router.Use(
		middleware.RequestID,
		httplog.RequestLogger(log.Logger),
	)

	router.Post("/createPost", s.handleCreatePost())
	router.Get("/getPosts", s.handleGetPosts())
	router.Post("/posts/{id}/like", s.handleLikePost())
	router.Post("/posts/{id}/unlike", s.handleUnlikePost())
	router.Post("/posts/{id}/comments", s.handleAddComment())

	return router
}

func (s *Server) Serve() {
	handler := s.GetHandler()
	err := http.ListenAndServe(fmt.Sprintf(":%d", s.cfg.ApplicationPort), handler)
	if err != nil {
		return
	}
}
