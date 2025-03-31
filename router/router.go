package router

import (
	"context"
	"errors"
	"fmt"
	"gitee.com/molonglove/goboot/gin"
	"go-demo/core"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var Engine *gin.Engine

func initRouter() {
	core.InitCore()

	fmt.Println("初始化路由")
}
func Run(port int64) {
	initRouter()

	server := &http.Server{
		Addr:           fmt.Sprintf("%d", port),
		Handler:        Engine,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: core.Config.Web.MaxHeaderBytes * 1024 * 1024,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("启动服务失败: %v", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Printf("服务关闭中...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("服务关闭失败: %v", err.Error())
	}
	fmt.Printf("服务关闭成功")
}
