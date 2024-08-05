package admin

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"godemo.com/demo/internal/code"
	"godemo.com/demo/internal/model"
	"godemo.com/demo/internal/util"
)

type ReqLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
}

var tokens = map[string]Token{
	"admin":  {Token: "admin-token"},
	"editor": {Token: "editor-token"},
}

type TokenInfo struct {
	Roles        []string `json:"roles"`
	Introduction string   `json:"introduction"`
	Avatar       string   `json:"avatar"`
	Name         string   `json:"name"`
}

var users = map[string]TokenInfo{
	"admin":  {Roles: []string{"admin"}, Introduction: "I am a super administrator", Avatar: "URL_ADDRESS", Name: "Super Admin"},
	"editor": {Roles: []string{"editor"}, Introduction: "I am an editor", Avatar: "URL_ADDRESS", Name: "Normal Editor"},
}

func Login(c *gin.Context) {
	var r ReqLogin
	err := c.ShouldBindBodyWithJSON(&r)
	if err != nil {
		util.Error(c, code.PARAMS_ERROR)
		return
	}
	fmt.Println(r)
	var admin model.Admin
	model.DB.Where("username=? and password=?", r.Username, r.Password).First(&admin)
	if admin.Id == 0 {
		util.Error(c, code.ADMIN_NOT_EXIST)
		return
	}
	fmt.Println(admin)
	util.Success(c, gin.H{
		"token": "admin",
	})
}

type ReqInfo struct {
	Token string `form:"token" json:"token"`
}

func Info(c *gin.Context) {
	var r ReqInfo
	err := c.ShouldBindQuery(&r)
	if err != nil {
		util.Error(c, code.PARAMS_ERROR)
		return
	}
	token := r.Token
	info := users[token]
	util.Success(c, info)
}

func Logout(c *gin.Context) {
	util.Success(c, nil)
}
