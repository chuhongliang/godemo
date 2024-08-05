package index

import (
	"github.com/gin-gonic/gin"
	"godemo.com/demo/internal/util"
)

func Index(c *gin.Context) {
	util.Success(c, "hello gin")
}
