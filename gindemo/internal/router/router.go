package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"godemo.com/demo/internal/api/admin"
	"godemo.com/demo/internal/api/auth"
	"godemo.com/demo/internal/api/index"
	"godemo.com/demo/internal/api/user"
	"godemo.com/demo/internal/middleware"
)

func InitRouter(r *gin.Engine) {
	// public api
	r.GET("/", index.Index)
	r.GET("/ping", index.Ping)

	authRouter := r.Group("/auth")
	authRouter.POST("/login", auth.Login)
	authRouter.POST("/register", auth.Register)

	adminRouter := r.Group("/admins")
	adminRouter.POST("/login", admin.Login)
	adminRouter.GET("/info", admin.Info)
	adminRouter.POST("/register", admin.Register)

	// seed
	adminRouter.POST("/seed/list", admin.GetSeedList)
	adminRouter.POST("seed/create", admin.CreateSeedItem)
	adminRouter.POST("seed/edit", admin.EditSeedItem)
	// task
	adminRouter.GET("/task/list", admin.GetTaskList)
	adminRouter.POST("/task/create", admin.CreateTaskItem)
	adminRouter.POST("/task/edit", admin.EditTaskItem)
	// user
	adminRouter.GET("/user/list", admin.GetUserList)
	adminRouter.GET("/user/info", admin.GetUserInfo)
	// shop
	adminRouter.POST("/shop/list", admin.GetShopList)
	adminRouter.POST("/shop/create", admin.CreateShopItem)
	adminRouter.POST("/shop/edit", admin.EditShopItem)

	// protected api
	v1 := r.Group("/api/v1")
	v1.Use(middleware.JwtAuthMiddleware())

	userRouter := v1.Group("/user")
	userRouter.GET("/", user.GetUser)

	fmt.Println("init router success")
}
