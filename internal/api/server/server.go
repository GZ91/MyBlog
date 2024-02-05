package server

import (
	"errors"
	"net/http"

	"github.com/GZ91/MyBlog/internal/api/handlers"
	"github.com/GZ91/MyBlog/internal/api/middleware"
	"github.com/GZ91/MyBlog/internal/service"
	"github.com/GZ91/MyBlog/internal/storage"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type Server struct {
	http.Server
	Logger *zap.Logger
}

func New() *Server {
	return &Server{}
}

func (s *Server) Configure(logger *zap.Logger) {
	s.Logger = logger
}

func (s *Server) Start() error {
	NodeStorage := storage.New(s.Logger)
	NodeService := service.New(s.Logger, NodeStorage)
	s.Addr = ":8080"
	s.Handler = s.routing(handlers.New(s.Logger, NodeService))

	if err := s.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			s.Logger.Error("server startup error", zap.Error(err))
		}
	}
	return nil
}

func (s *Server) routing(handls *handlers.Handlers) *chi.Mux {
	router := chi.NewRouter()
	middleware := middleware.New(s.Logger)

	//router.Use(middleware.WithLogging)
	router.Use(middleware.Compress)
	router.Use(middleware.Authentication)
	router.Use(middleware.CalculateSize)

	fs := http.FileServer(http.Dir("../../source/"))
	router.Handle("/source/*", http.StripPrefix("/source/", fs))
	router.Get("/", handls.MainPage)

	return router
}
