package service

import (
	"github.com/GZ91/MyBlog/internal/models"
	"go.uber.org/zap"
)

type NodeStorager interface {
	Authorized(userID string) (bool, error)
	Login(login, password, userID string) (bool, error)
	GetArts() ([]models.Art, error)
}

type Service struct {
	NodeStorage NodeStorager
	logger      *zap.Logger
}

func New(logger *zap.Logger, NodeS NodeStorager) *Service {
	return &Service{
		NodeStorage: NodeS,
		logger:      logger,
	}
}
