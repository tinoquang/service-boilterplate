package handler

import "go.uber.org/zap"

type APIHandler struct {
	defaultLogger *zap.Logger
}

func New(l *zap.Logger) *APIHandler {
	return &APIHandler{
		defaultLogger: l,
	}
}
