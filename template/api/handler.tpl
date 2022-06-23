package {{.PkgName}}

import (
	"net/http"

	"github.com/sprchu/tiktok/servebase"
	"github.com/sprchu/tiktok/servebase/errno"
	{{.ImportPackages}}

	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
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

		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
		if err != nil {
			{{if not .HasResp}}resp := servebase.CommonResponse{}
			{{end}}resp.StatusCode = err.(errno.ErrNo).ErrCode
			resp.StatusMsg = err.(errno.ErrNo).ErrMsg
			{{if .HasResp}}httpx.OkJson(w, resp){{else}}httpx.Ok(w){{end}}
		} else {
			{{if not .HasResp}}resp := servebase.CommonResponse{}
			{{end}}resp.StatusCode = errno.SuccessCode
			{{if .HasResp}}httpx.OkJson(w, resp){{else}}httpx.Ok(w){{end}}
		}
	}
}
