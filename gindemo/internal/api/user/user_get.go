package user

import (
	"github.com/gin-gonic/gin"
	"godemo.com/demo/internal/code"
	"godemo.com/demo/internal/model"
	"godemo.com/demo/internal/util"
)

func GetUser(c *gin.Context) {
	userId, _ := c.Get("user_id")

	var user model.User
	model.DB.Where("id =?", userId).First(&user)

	if user.Id == 0 {
		util.Error(c, code.USER_NOT_EXIST)
		return
	}

	var lands []model.Land
	model.DB.Where("user_id =?", userId).Find(&lands)

	var items []model.Item
	model.DB.Where("user_id =?", userId).Find(&items)

	util.Success(c, gin.H{
		"user":  user,
		"lands": lands,
		"items": items,
	})
}
