package service

import "go.uber.org/zap"

type NodeStorager interface {
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
