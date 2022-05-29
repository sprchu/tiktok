package middleware

import (
	"context"
	"net/http"

	"github.com/ByteDance-camp/TickTalk/servebase"
	"github.com/ByteDance-camp/TickTalk/servebase/errno"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type AuthMiddleware struct {
	secret string
}

type UID string

func NewAuthMiddleware(st string) *AuthMiddleware {
	return &AuthMiddleware{secret: st}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.FormValue("token")
		uid, err := servebase.ParseToken(token, m.secret)
		if err != nil {
			httpx.OkJson(w, servebase.CommonResponse{
				StatusCode: errno.AuthErrCode,
				StatusMsg:  err.Error(),
			})
			return
		}
		r = r.WithContext(context.WithValue(r.Context(), UID("uid"), uid))

		next(w, r)
	}
}
