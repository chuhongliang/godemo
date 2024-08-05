package index

import (
	"github.com/gin-gonic/gin"
	"godemo.com/demo/internal/util"
)

func Ping(c *gin.Context) {
	util.Success(c, "pong")
}
