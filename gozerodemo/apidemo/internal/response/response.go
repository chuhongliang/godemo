package response

import (
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/rest/httpx"
	Code "gozerodemo.com/apidemo/internal/code"
	"gozerodemo.com/apidemo/internal/custom"
)

type SuccessBody struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
	Now     time.Time   `json:"now"`
}

type FailBody struct {
	Success bool      `json:"success"`
	Code    int       `json:"code"`
	Msg     string    `json:"msg"`
	Now     time.Time `json:"now"`
}

func Success(w http.ResponseWriter, resp interface{}) {
	var body SuccessBody
	body.Success = true
	body.Code = 200
	body.Msg = "OK"
	body.Data = resp
	body.Now = time.Now()
	httpx.OkJson(w, body)
}

func Fail(w http.ResponseWriter, err custom.CustomError) {
	var body FailBody
	body.Success = false
	body.Code = err.Code
	body.Msg = Code.Text(err.Code)
	body.Now = time.Now()
	httpx.OkJson(w, body)
}
