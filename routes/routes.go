package routes

import (
	"graduation/controllers"
	"graduation/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

// gin.Context 是一个非常重要的结构体，
// 它承载着 HTTP 请求的上下文信息，
// 包含了请求、响应、路由、以及与请求生命周期相关的一些方法。
func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}

	r := gin.New()
	r.Use(CORSMiddleware())

	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := r.Group("/api/v1").Use(controllers.TimeoutMiddleware())
	// 注册业务路由
	v1.POST("/signup", controllers.SignupHandler)
	v1.POST("/login", controllers.LoginHandler)

	v1.POST("/emotional", controllers.EmotionalHandler)
	v1.GET("/emotional/count/:user_id", controllers.EmotionalCountHandler)

	v1.POST("/checkin/:user_id", controllers.CheckinHandler)

	// v1.Use(middlewares.JWTAuthMiddleware()) // 应用JWT认证中间件



	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	
	return r
}
