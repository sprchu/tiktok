package middleware

import (
	"net/http"

	"github.com/sprchu/tiktok/servebase"
	"github.com/sprchu/tiktok/servebase/errno"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type AuthMiddleware struct {
	secret string
}

func NewAuthMiddleware(secret string) *AuthMiddleware {
	return &AuthMiddleware{secret: secret}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.URL.Query().Get("token")
		_, err := servebase.ParseToken(token, m.secret)
		if err != nil {
			httpx.OkJson(w, servebase.CommonResponse{
				StatusCode: errno.AuthErrCode,
				StatusMsg:  err.Error(),
			})
			return
		}

		next(w, r)
	}
}
