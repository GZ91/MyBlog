package storage

import "go.uber.org/zap"

type Storage struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) *Storage {
	return &Storage{
		logger: logger,
	}
}
