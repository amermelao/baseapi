package handler

import (
	"baseapi/api"
	"context"
	"log/slog"

	"time"
)

type handler struct {
	logger  *slog.Logger
	storage dataStorer
}

type dataStorer interface {
}

func NewHandler(logger *slog.Logger, p dataStorer) api.StrictServerInterface {
	return handler{logger: logger, storage: p}
}

func (handler) GetAlive(
	ctx context.Context,
	request api.GetAliveRequestObject,
) (api.GetAliveResponseObject, error) {
	ok := api.GetAlive200JSONResponse{
		Status:    "ok",
		Timestamp: time.Now(),
	}
	return ok, nil
}
