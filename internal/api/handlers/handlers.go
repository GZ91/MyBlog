package handlers

import "go.uber.org/zap"

type NodeServicer interface{}

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
