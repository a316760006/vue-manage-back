package routes

import (
	"net/http"
	"vue-manage-back/app/controllers/app"
	"vue-manage-back/app/controllers/common"
	"vue-manage-back/app/middleware"
	"vue-manage-back/app/services"

	"github.com/gin-gonic/gin"
)

// SetApiGroupRoutes 定义 api 分组路由
func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	// router.GET("/test", func(c *gin.Context) {
	// 	time.Sleep(5 * time.Second)
	// 	c.String(http.StatusOK, "success")
	// })
	// router.POST("/user/register", func(c *gin.Context) {
	// 	var form request.Register
	// 	if err := c.ShouldBindJSON(&form); err != nil {
	// 		c.JSON(http.StatusOK, gin.H{
	// 			"error": request.GetErrorMsg(form, err),
	// 		})
	// 		return
	// 	}
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "success",
	// 	})
	// })

	// 注册
	router.POST("/auth/register", app.Register)
	// 登录
	router.POST("/auth/login", app.Login)
	// token 鉴权
	authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	{
		authRouter.POST("/auth/info", app.Info)
		authRouter.POST("/auth/logout", app.Logout)
		authRouter.POST("/image_upload", common.ImageUpload)
	}
	router.GET("/ws", app.Ws)
}
