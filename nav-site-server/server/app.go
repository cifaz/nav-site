package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	conf "nav-site-server/config"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"
)

func initLog() {
	log.SetFormatter(&log.JSONFormatter{}) //设置日志的输出格式为json格式，还可以设置为text格式
	log.SetOutput(os.Stdout)               //设置日志的输出为标准输出
	log.SetLevel(log.InfoLevel)            //设置日志的显示级别，这一级别以及更高级别的日志信息将会输出
}

func init() {
	//以package级别方式使用日志
	initLog()

	// 读取参数
	var confDir string
	for _, item := range os.Args {
		itemArr := strings.Split(item, "=")
		if itemArr[0] == "-conf-dir" || itemArr[0] == "conf-dir" {
			confDir = itemArr[1]
		}
	}

	s, err := conf.InitConfig(confDir)
	if err != nil {
		log.Info("init server config failed, error: ", err)
		//log.Println("init server config failed, error: ", err)
		os.Exit(1)
	}
	conf.App = *s
	//log.Info("init config:")
	log.Info("init config:", conf.App)
	//log.Printf("init config:")
	//log.Printf("%+v", conf.App)
}

func Run() {
	engine := gin.Default()
	defer func() {
		if err := conf.App.Store.FileSync.CloseStoreFile(); err != nil {
			log.Info("close store file resource failed, error :", err)
			//log.Println("close store file resource failed, error :", err)
			os.Exit(1)
		}
	}()

	InitHtmlResource(engine)

	// 读取本地静态资源
	//InitLocalResource(engine)
	// embed方式读取静态资源
	InitEmbedResource(engine)

	Router(engine)
	engine.NoMethod(conf.HandleNotFound)
	engine.NoRoute(conf.HandleNotFound)

	port := fmt.Sprintf(":%s", conf.App.Server.Port)
	//log.Println("listen port", port)
	log.Info("listen port", port)

	srv := &http.Server{
		Addr:    port,
		Handler: engine,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			//log.Fatalf("listen: %s\n", err)
			log.Fatal("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Info("Shutdown Server ...")
	//log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		//log.Fatal("Server Shutdown:", err)
		log.Fatal("Server Shutdown:", err)
	}
	log.Info("Server exiting")
	//log.Println("Server exiting")

}
