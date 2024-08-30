package main

import (
	"apidoc/config"
	"apidoc/environment"
	"apidoc/middleware"
	"apidoc/route"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	env  *environment.Env
	conf *config.Config
)

func main() {
	// 加载常用的变量
	env = environment.Load()
	// 加载配置文件
	conf = config.Load(env.CurDir + env.Separate + "config.yaml")
	// todo 定义日志格式，载入自定义日志目录和名称

	// 新建路由
	r := gin.New()
	// 添加中间件
	middleware.Load(r)
	// 加载路由
	route.Load(r)

	// 服务器启动
	port, portErr := conf.GetString("server.port")
	if portErr != nil {
		log.Fatal(portErr)
	}
	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if shutdownErr := srv.Shutdown(ctx); shutdownErr != nil {
		log.Fatal("Server Shutdown:", shutdownErr)
	}
	log.Println("Server exiting")
}
