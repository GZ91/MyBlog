package handlers

import (
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
}

func New(logger *zap.Logger, NodeS NodeServicer) *Handlers {
	return &Handlers{
		NodeService: NodeS,
		logger:      logger,
	}
}

type Page struct {
	Content    string
	AlterLabel string
}
