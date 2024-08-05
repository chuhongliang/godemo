package admin

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"godemo.com/demo/internal/code"
	"godemo.com/demo/internal/model"
	"godemo.com/demo/internal/util"
)

func GetUserList(c *gin.Context) {
	var r ReqList
	err := c.ShouldBindJSON(&r)
	if err != nil {
		util.Error(c, code.PARAMS_ERROR)
		return
	}

	var list []model.User
	offset := (r.Page - 1) * r.PageSize
	model.DB.Limit(r.PageSize).Offset(offset).Find(&list)

	var count int64
	model.DB.Model(&model.User{}).Count(&count)

	totalPage := 1
	if count != 0 {
		fmt.Println("count:", count)
		totalPage = int(count / int64(r.PageSize))
	}
	fmt.Println("totalPage:", totalPage)
	if totalPage <= 0 {
		totalPage = 1
	}
	util.Success(c, gin.H{
		"items":      list,
		"page":       r.Page,
		"page_size":  r.PageSize,
		"total_page": totalPage,
	})
}

type ReqGetUserInfo struct {
	UserId int `json:"user_id" binding:"required"`
}

func GetUserInfo(c *gin.Context) {
	var r ReqGetUserInfo
	err := c.ShouldBindBodyWithJSON(&r)
	if err != nil {
		util.Error(c, code.PARAMS_ERROR)
		return
	}
	var user model.User
	model.DB.First(&user, r.UserId)
	util.Success(c, gin.H{
		"user_info": user,
	})
}
