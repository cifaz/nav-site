package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	conf "nav-site-server/config"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func init() {
	s, err := conf.InitConfig()
	if err != nil {
		log.Println("init server config failed, error: ", err)
		os.Exit(1)
	}
	conf.App = *s
	log.Printf("init config:")
	log.Printf("%+v", conf.App)
}

func Run() {
	engine := gin.Default()
	defer func() {
		if err := conf.App.Store.FileSync.CloseStoreFile(); err != nil {
			log.Println("close store file resource failed, error :", err)
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
	log.Println("listen port", port)

	srv := &http.Server{
		Addr:    port,
		Handler: engine,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

}
