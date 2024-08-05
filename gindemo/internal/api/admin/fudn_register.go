package admin

import (
	"github.com/gin-gonic/gin"
	"godemo.com/demo/internal/code"
	"godemo.com/demo/internal/model"
	"godemo.com/demo/internal/util"
)

type ReqCreateAdmin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// @Summary 注册管理员
// @Produce json
// @Param Username query string false "文章名称"
// @Param Password query int false "标签ID"
// @Success 200 {object} model.Admin "成功"
// @Router /admin/register [post]
func Register(c *gin.Context) {
	var r ReqCreateAdmin
	err := c.ShouldBindBodyWithJSON(&r)
	if err != nil {
		util.Error(c, code.PARAMS_ERROR)
		return
	}
	admin := model.Admin{
		Username: r.Username,
		Password: r.Password,
	}
	model.DB.Create(&admin)
	if admin.Id == 0 {
		util.Error(c, code.USER_EXIST)
		return
	}
	util.Success(c, admin)
}
