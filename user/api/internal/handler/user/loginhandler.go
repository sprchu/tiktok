package user

import (
	"net/http"

	"github.com/ByteDance-camp/TickTalk/servebase"
	"github.com/ByteDance-camp/TickTalk/servebase/errno"
	"github.com/ByteDance-camp/TickTalk/user/api/internal/logic/user"
	"github.com/ByteDance-camp/TickTalk/user/api/internal/svc"
	"github.com/ByteDance-camp/TickTalk/user/api/internal/types"

	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
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

		l := user.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
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
