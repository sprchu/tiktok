package video

import (
	"net/http"

	"github.com/ByteDance-camp/TickTalk/servebase"
	"github.com/ByteDance-camp/TickTalk/servebase/errno"
	"github.com/ByteDance-camp/TickTalk/videomgr/api/internal/logic/video"
	"github.com/ByteDance-camp/TickTalk/videomgr/api/internal/middleware"
	"github.com/ByteDance-camp/TickTalk/videomgr/api/internal/svc"
	"github.com/ByteDance-camp/TickTalk/videomgr/api/internal/types"

	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func PublishListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PublishListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, servebase.CommonResponse{
				StatusCode: errno.ParamErrCode,
				StatusMsg:  err.Error(),
			})
			return
		}
		if err := validator.New().StructCtx(r.Context(), req); err != nil {
			httpx.OkJson(w, servebase.CommonResponse{
				StatusCode: errno.ParamErrCode,
				StatusMsg:  err.Error(),
			})
			return
		}

		uid, ok := r.Context().Value(middleware.UID("uid")).(int64)
		if !ok || uid != req.UserID {
			httpx.OkJson(w, servebase.CommonResponse{
				StatusCode: errno.AuthErrCode,
				StatusMsg:  "failed to validate uid",
			})
			return
		}

		l := video.NewPublishListLogic(r.Context(), svcCtx)
		resp, err := l.PublishList(&req)
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
