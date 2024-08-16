package auth

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gozerodemo.com/apidemo/internal/code"
	"gozerodemo.com/apidemo/internal/custom"
	"gozerodemo.com/apidemo/internal/logic/auth"
	"gozerodemo.com/apidemo/internal/response"
	"gozerodemo.com/apidemo/internal/svc"
	"gozerodemo.com/apidemo/internal/types"
)

func PingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PingReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := auth.NewPingLogic(r.Context(), svcCtx)
		resp, err := l.Ping(&req)
		if err != nil {
			if customErr, ok := err.(custom.CustomError); ok {
				response.Fail(w, customErr)
			} else {
				response.Fail(w, custom.CustomError{
					Code: code.ERROR,
				})
			}
		} else {
			response.Success(w, resp)
		}
	}
}