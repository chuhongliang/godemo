package main

import (
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "godemo.com/demo/docs"
	"godemo.com/demo/internal/cron"
	"godemo.com/demo/internal/model"
	"godemo.com/demo/internal/router"
	"godemo.com/demo/internal/util"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file. %v\n", err)
	}
}

// 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

func main() {
	// 初始化配置
	util.InitConfig()
	// 初始化MYSQL
	model.InitGormClient()
	// 初始化定时器
	cron.InitCron()

	r := gin.Default()
	r.Use(Cors())
	r.Static("/assets", "./assets")
	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.POST("admins/upload", func(c *gin.Context) {
		// 单文件
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		dst := "./assets/img/" + file.Filename
		// 上传文件至指定的完整文件路径
		c.SaveUploadedFile(file, dst)

		baseurl := "http://192.168.1.234:7001/assets/img"
		imgurl := baseurl + "/" + file.Filename

		util.Success(c, gin.H{
			"imgurl": imgurl,
		})
	})
	// 初始化路由
	router.InitRouter(r)

	r.Run(":7001") // 监听并在 0.0.0.0:7001 上启动服务
}
