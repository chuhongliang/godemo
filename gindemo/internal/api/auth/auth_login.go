package auth

import (
	"github.com/gin-gonic/gin"
	"godemo.com/demo/internal/code"
	"godemo.com/demo/internal/model"
	"godemo.com/demo/internal/token"
	"godemo.com/demo/internal/util"
)

type ReqLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var r ReqLogin
	err := c.ShouldBindBodyWithJSON(&r)
	if err != nil {
		util.Error(c, code.PARAMS_ERROR)
		return
	}

	var user model.User
	model.DB.Where("username =?", r.Username).First(&user)
	if user.Id == 0 {
		util.Error(c, code.USER_NOT_EXIST)
		return
	}

	userToken, _ := token.GenerateToken(uint(user.Id))
	util.Success(c, gin.H{
		"token": userToken,
	})

}
