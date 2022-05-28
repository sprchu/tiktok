package storage

import (
	"context"
	"mime/multipart"
)

var handler FileHandler = &nilHandler{}

type nilHandler struct{}

type FileHandler interface {
	Upload(ctx context.Context, file *multipart.FileHeader) (url string, err error)
}

func RegisterHandler(hd FileHandler) {
	handler = hd
}

func Upload(ctx context.Context, file *multipart.FileHeader) (url string, err error) {
	return handler.Upload(ctx, file)
}

func (h *nilHandler) Upload(ctx context.Context, file *multipart.FileHeader) (string, error) {
	return "create by nilHandler", nil
}
