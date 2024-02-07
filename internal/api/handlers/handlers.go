package handlers

import (
	"github.com/GZ91/MyBlog/internal/app/config"
	"github.com/GZ91/MyBlog/internal/models"
	"go.uber.org/zap"
)

type NodeServicer interface {
	Authorized(userID string) (bool, error)
	Login(login, password, userID string) (bool, error)
	GetArts() ([]models.Art, error)
}

type Handlers struct {
	NodeService NodeServicer
	logger      *zap.Logger
	config      *config.Config
}

func New(logger *zap.Logger, NodeS NodeServicer, config *config.Config) *Handlers {
	return &Handlers{
		NodeService: NodeS,
		logger:      logger,
		config:      config,
	}
}
