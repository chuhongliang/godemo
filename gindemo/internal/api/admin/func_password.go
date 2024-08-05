package admin

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"godemo.com/demo/internal/code"
	"godemo.com/demo/internal/model"
	"godemo.com/demo/internal/util"
)

type ReqParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func ChangePassword(c *gin.Context) {
	var r ReqParams
	err := c.ShouldBindBodyWithJSON(&r)
	if err != nil {
		util.Error(c, code.PARAMS_ERROR)
		return
	}
	fmt.Println(r)
	model.DB.Model(model.Admin{}).Where(&model.Admin{Username: r.Username}).Update("password", r.Password)
	util.Success(c, "修改成功")
}
