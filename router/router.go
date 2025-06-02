package router

import (
	"context"
	"errors"
	"fmt"
	"gitee.com/molonglove/goboot/gin"
	"go-demo/app/controller"
	"go-demo/core"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var Engine *gin.Engine

func initRouter() {
	core.InitCore()
	gin.SetMode(core.Config.Web.RunModel)
	core.Log.Info("初始化配置文件成功")
	Engine = gin.Default(core.Log)
	Engine.Use(CorsMiddle()) //跨域
	//Engine.Use(JwtMiddle())  //jwt
	root := Engine.Group(core.Config.Web.ContextPath)
	{
		user := root.Group("user")
		{
			user.GET("captchaImage", controller.User.CaptchaImage) //获取验证码
			user.POST("login", controller.User.Login)              //登录
			user.GET("GetUserInfo", controller.User.GetUserInfo)   //用户信息
			user.POST("page", controller.User.Page)                //分页

		}
	}

}
func Run(port int64) {
	initRouter()
	core.Log.Info("服务启动在端口[%d]", port)
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        Engine,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: core.Config.Web.MaxHeaderBytes * 1024 * 1024,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			core.Log.Error("启动服务失败: %s\n", err.Error())
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	core.Log.Info("服务正在停止...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		core.Log.Error("服务停止失败: %v", err.Error())
	}
	core.Log.Info("服务停止...")
}
