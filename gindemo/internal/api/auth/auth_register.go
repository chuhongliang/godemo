package auth

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	Code "godemo.com/demo/internal/code"
	"godemo.com/demo/internal/model"
	"godemo.com/demo/internal/util"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var r RegisterRequest
	err := c.ShouldBindBodyWithJSON(&r)
	if err != nil {
		fmt.Println("register failed")
		util.Error(c, Code.PARAMS_ERROR)
		return
	}

	var user model.User
	model.DB.Where("username =?", r.Username).First(&user)
	fmt.Println("user", user)
	if user.Id != 0 {
		util.Error(c, Code.USER_EXIST)
		return
	}

	extra := model.Extra{
		Guide: 0,
	}
	extraStr, err := json.Marshal(extra)
	if err != nil {
		util.Error(c, Code.PARAMS_ERROR)
	}

	user = model.User{
		Username: r.Username,
		Password: r.Password,
		Extra:    string(extraStr),
	}
	model.DB.Create(&user)
	if user.Id == 0 {
		util.Error(c, Code.USER_EXIST)
		return
	}

	fmt.Println(&user)

	for x := 1; x <= 12; x++ {
		land := model.Land{
			UserId:   user.Id,
			Position: x,
		}
		if x <= 3 {
			land.Isunlock = true
		}
		model.DB.Create(&land)
	}

	var seed model.Seed
	model.DB.Where("id =?", 1).First(&seed)
	if seed.Id != 0 {
		item := model.Item{
			UserId:     user.Id,
			ItemType:   1,
			ItemId:     seed.Id,
			ItemTitle:  seed.Title,
			ItemNumber: 3,
			ItemImgurl: seed.Imgurl,
			ItemDesc:   seed.Desc,
		}
		model.DB.Create(&item)
	}

	var userInfo model.User
	model.DB.Where("id =?", user.Id).First(&userInfo)

	util.Success(c, gin.H{
		"user": userInfo,
	})
}
