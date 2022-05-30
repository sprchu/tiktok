package storage

import (
	"bytes"
	"context"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/zeromicro/go-zero/core/logx"
)

type localHandler struct {
	path string
}

func NewLocalHandler(addr, path string) (*localHandler, chan struct{}) {
	stop := make(chan struct{})
	h := http.FileServer(http.Dir(path))
	go func() {
		err := http.ListenAndServe(addr, h)
		if err != nil {
			logx.Errorf("file server error: %v", err)
			stop <- struct{}{}
		}
	}()
	return &localHandler{path: path}, stop
}

func (h *localHandler) Upload(ctx context.Context, file *multipart.FileHeader) (string, error) {
	f, err := file.Open()
	if err != nil {
		logx.Errorf("open multipart file error: %w", err)
		return "", err
	}
	defer f.Close()

	content, err := ioutil.ReadAll(f)
	if err != nil {
		logx.Errorf("read multipart file error: %w", err)
		return "", err
	}

	sha := sha256.Sum256(content)
	storeFileName := fmt.Sprintf("%s/%x", h.path, sha)
	_, err = os.Stat(storeFileName)
	if err == nil {
		return storeFileName, nil
	}
	if !os.IsExist(err) {
		logx.Errorf("failed to read file info: %w, sha256: %x", err, sha)
		return "", err
	}

	storeFile, err := os.Open(storeFileName)
	if err != nil {
		logx.Errorf("failed to open file to store: %w, sha256: %x", err, sha)
		return "", err
	}
	defer storeFile.Close()
	_, err = io.Copy(storeFile, bytes.NewBuffer(content))
	if err != nil {
		logx.Errorf("write file error: %v", err)
		return "", err
	}

	return fmt.Sprintf("%s/%x", h.path, sha), nil
}
