package middleware

import "go.uber.org/zap"

type Middleware struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) *Middleware {
	return &Middleware{
		logger: logger,
	}
}
