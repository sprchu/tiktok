package storage

import (
	"bytes"
	"context"
	"crypto/sha256"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

const scale = "600x800"

type localHandler struct {
	addr string
	path string
}

func NewLocalHandler(addr, path string) (*localHandler, chan struct{}) {
	stop := make(chan struct{})
	h := http.FileServer(http.Dir(path))
	go func() {
		err := http.ListenAndServe(addr[strings.Index(addr, ":"):], h)
		if err != nil {
			logx.Errorf("file server error: %v", err)
			stop <- struct{}{}
		}
	}()
	return &localHandler{addr: addr, path: path}, stop
}

func (h *localHandler) Upload(ctx context.Context, file *multipart.FileHeader) (string, string, error) {
	f, err := file.Open()
	if err != nil {
		logx.Errorf("open multipart file error: %w", err)
		return "", "", err
	}
	defer f.Close()

	content, err := ioutil.ReadAll(f)
	if err != nil {
		logx.Errorf("read multipart file error: %w", err)
		return "", "", err
	}

	sha := sha256.Sum256(content)
	storeFileName := fmt.Sprintf("%s/%x", h.path, sha)
	ok := fs.ValidPath(storeFileName)
	if ok {
		return storeFileName, storeFileName + ".png", nil
	}

	storeFile, err := os.Create(storeFileName)
	if err != nil {
		logx.Errorf("failed to create file: %w, sha256: %x", err, sha)
		return "", "", err
	}
	defer storeFile.Close()
	_, err = io.Copy(storeFile, bytes.NewBuffer(content))
	if err != nil {
		logx.Errorf("write file error: %v", err)
		return "", "", err
	}

	err = h.genCover(storeFileName)
	if err != nil {
		logx.Errorf("gen cover error: %v", err)
		return "", "https://xcj-pic.oss-cn-beijing.aliyuncs.com/kedaya.jpg", err
	}

	return fmt.Sprintf("%s/%x", h.addr, sha), storeFileName + ".png", nil
}

func (h *localHandler) genCover(filename string) error {
	command := fmt.Sprintf(
		"ffmpeg -ss 0.5 -i %s -vframes 1 -s %s -f image2 %s",
		filename,
		scale,
		filename+".png",
	)
	cmd := exec.Command(command)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
