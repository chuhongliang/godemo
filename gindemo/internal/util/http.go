package util

import (
	"time"

	"github.com/gin-gonic/gin"
	Code "godemo.com/demo/internal/code"
)

func Success(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"success": true,
		"code":    Code.SUCCESS,
		"data":    data,
		"now":     time.Now().Unix(),
	})
}

func Error(c *gin.Context, code int) {
	c.JSON(200, gin.H{
		"success": false,
		"code":    code,
		"msg":     Code.Text(code),
		"now":     time.Now().Unix(),
	})
}
