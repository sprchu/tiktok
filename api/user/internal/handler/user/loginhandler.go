package user

import (
	"net/http"

	"github.com/sprchu/tiktok/api/user/internal/logic/user"
	"github.com/sprchu/tiktok/api/user/internal/svc"
	"github.com/sprchu/tiktok/api/user/internal/types"
	"github.com/sprchu/tiktok/servebase"
	"github.com/sprchu/tiktok/servebase/errno"

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
