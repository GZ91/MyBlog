package server

import (
	"context"
	"errors"
	"net/http"

	"github.com/GZ91/MyBlog/internal/api/handlers"
	"github.com/GZ91/MyBlog/internal/api/middleware"
	"github.com/GZ91/MyBlog/internal/app/config"
	"github.com/GZ91/MyBlog/internal/service"
	"github.com/GZ91/MyBlog/internal/storage"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type Server struct {
	http.Server
	Logger *zap.Logger
	Config *config.Config
}

func New() *Server {
	return &Server{}
}

func (s *Server) Configure(logger *zap.Logger, config *config.Config) {
	s.Logger = logger
	s.Config = config
}

func (s *Server) Start() error {
	ctx := context.Background()
	NodeStorage := storage.New(s.Logger, s.Config)
	if err := NodeStorage.Up(ctx); err != nil {
		s.Logger.Error("storage up error", zap.Error(err))
		return err
	}
	NodeService := service.New(s.Logger, NodeStorage)
	s.Addr = ":80"
	s.Handler = s.routing(handlers.New(s.Logger, NodeService, s.Config))

	if err := s.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			s.Logger.Error("server startup error", zap.Error(err))
		}
	}
	NodeStorage.Close()
	return nil
}

func (s *Server) routing(handls *handlers.Handlers) *chi.Mux {
	router := chi.NewRouter()
	middleware := middleware.New(s.Logger)

	//router.Use(middleware.WithLogging)
	router.Use(middleware.Compress)
	router.Use(middleware.Authentication)
	router.Use(middleware.CalculateSize)

	fs := http.FileServer(http.Dir(s.Config.GetMainPath() + "/source/"))
	router.Handle("/source/*", http.StripPrefix("/source/", fs))

	router.Get("/", handls.Index)
	router.Get("/login", handls.Login)
	router.Post("/login", handls.Login)
	router.Post("/login/post", handls.LoginPost)

	return router
}
