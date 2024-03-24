package main

import (
	"fmt"
	"log"
	"net/http"
	"story-api/models"
	"story-api/pkg/setting"
	"story-api/routers"

	"github.com/gin-gonic/gin"
	// "github.com/vitorbarbarisi/story-api/models"
	// "github.com/vitorbarbarisi/story-api/pkg/gredis"
	// "github.com/vitorbarbarisi/story-api/pkg/logging"
	// "github.com/vitorbarbarisi/story-api/pkg/util"
)

func init() {
	setting.Setup()
	models.Setup()
}

// @title Story API
// @version 1.0
// @description An example of gin
// @termsOfService https://github.com/vitorbarbarisi/story-api
func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()

	// If you want Graceful Restart, you need a Unix system and download github.com/fvbock/endless
	//endless.DefaultReadTimeOut = readTimeout
	//endless.DefaultWriteTimeOut = writeTimeout
	//endless.DefaultMaxHeaderBytes = maxHeaderBytes
	//server := endless.NewServer(endPoint, routersInit)
	//server.BeforeBegin = func(add string) {
	//	log.Printf("Actual pid is %d", syscall.Getpid())
	//}
	//
	//err := server.ListenAndServe()
	//if err != nil {
	//	log.Printf("Server err: %v", err)
	//}
}
