
# 初始化Handler模版
```bash 
package {{.PkgName}}

import (
        "net/http"

        "github.com/zeromicro/go-zero/rest/httpx"
        {{.ImportPackages}}
)

{{if .HasDoc}}{{.Doc}}{{end}}
func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
        return func(w http.ResponseWriter, r *http.Request) {
                {{if .HasRequest}}var req types.{{.RequestType}}
                if err := httpx.Parse(r, &req); err != nil {
                        httpx.ErrorCtx(r.Context(), w, err)
                        return
                }

                {{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
                {{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
                if err != nil {
                        httpx.ErrorCtx(r.Context(), w, err)
                } else {
                        {{if .HasResp}}httpx.OkJsonCtx(r.Context(), w, resp){{else}}httpx.Ok(w){{end}}
                }
        }
}
```

# 自定义Handler模版
```bash
package {{.PkgName}}

import (
    "net/http"
    "gozerodemo.com/apidemo/internal/response"
    "github.com/zeromicro/go-zero/rest/httpx"
    {{.ImportPackages}}
)

func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
     return func(w http.ResponseWriter, r *http.Request) {
                {{if .HasRequest}}var req types.{{.RequestType}}
                if err := httpx.Parse(r, &req); err != nil {
                        httpx.ErrorCtx(r.Context(), w, err)
                        return
                }

                {{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
                {{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
                if err != nil {
                        {{if .HasResp}}response.Fail(w, err){{end}}
                } else {
                        {{if .HasResp}}response.Success(w, resp){{end}}
                }
        }
}
```


