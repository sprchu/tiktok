package video

import (
	"net/http"

	"github.com/sprchu/tiktok/api/videomgr/internal/logic/video"
	"github.com/sprchu/tiktok/api/videomgr/internal/svc"
	"github.com/sprchu/tiktok/api/videomgr/internal/types"
	"github.com/sprchu/tiktok/api/videomgr/storage"
	"github.com/sprchu/tiktok/servebase"
	"github.com/sprchu/tiktok/servebase/errno"

	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

const maxFileSize = 1 << 20 * 200

func PublishActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PublishActionRequest
		sendErr := func(code, msg string) {
			httpx.OkJson(w, servebase.CommonResponse{
				StatusCode: code,
				StatusMsg:  msg,
			})
		}
		if err := httpx.Parse(r, &req); err != nil {
			sendErr(errno.ParamErrCode, err.Error())
			return
		}
		if err := validator.New().StructCtx(r.Context(), req); err != nil {
			sendErr(errno.ParamErrCode, err.Error())
			return
		}

		if err := r.ParseMultipartForm(maxFileSize); err != nil {
			sendErr(errno.ParamErrCode, err.Error())
			return
		}

		files := r.MultipartForm.File["data"]
		if len(files) != 1 {
			sendErr(errno.UploadErrCode, "file count must be 1")
			return
		}

		url, cover, err := storage.Upload(r.Context(), files[0])
		if err != nil {
			sendErr(errno.UploadErrCode, err.Error())
			return
		}
		logx.Infof("upload file at: %s", url)

		req.Url = url
		req.Cover = cover

		l := video.NewPublishActionLogic(r.Context(), svcCtx)
		resp, err := l.PublishAction(&req)
		if err != nil {
			resp.StatusCode = err.(errno.ErrNo).ErrCode
			resp.StatusMsg = err.(errno.ErrNo).ErrMsg
			httpx.OkJson(w, resp)
		} else {
			resp.StatusCode = errno.SuccessCode
			httpx.OkJson(w, resp)
		}
	}
}
