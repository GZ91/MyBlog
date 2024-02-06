package app

import (
	"github.com/GZ91/MyBlog/internal/api/server"
	"github.com/GZ91/MyBlog/internal/app/config"
	"github.com/GZ91/MyBlog/internal/app/logger"
	"go.uber.org/zap"
)

type App struct {
	Logger *zap.Logger
	Config *config.Config
}

func New(lvlLogger string, config *config.Config) *App {
	logger, err := logger.Initializing(lvlLogger)
	if err != nil {
		panic(err)
	}
	return &App{
		Logger: logger,
		Config: config,
	}
}

func (a *App) Start() error {
	server := server.New()
	server.Configure(a.Logger, a.Config)
	err := server.Start()
	if err != nil {
		return err
	}
	return nil
}
