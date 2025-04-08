package {{.PkgName}}

import (
	"net/http"

	{{.ImportPackages}}

	"github.com/lerryxu/go-zero/common/xerr"
)

func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
		if err := xerr.Parse(r, &req); err != nil {
			xerr.Response(r.Context(), w, nil, xerr.NewError(xerr.REUQEST_PARAM_ERROR, err.Error()))
			return
		}

		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r, w, svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
		{{if .HasResp}}xerr.Response(r.Context(), w, resp, err){{else}}xerr.Response(r.Context(), w, nil, err){{end}}
	}
}
