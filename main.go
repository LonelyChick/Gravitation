package main

import (
	"fmt"
	"github.com/gin-gonic/gin"

	//"github.com/gin-gonic/gin"
	"net/http"


	"gravitation/pkg/setting"
	"gravitation/routers"
)

func init() {
	setting.Setup()
}

func main() {
	gin.SetMode(setting.ServerSetting.Runmode)

	routerInit := routers.Router()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	litsenPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeadBytes := 1 << 20

	server := &http.Server{
		Addr:              litsenPoint,
		Handler:           routerInit,
		ReadTimeout:       readTimeout,
		WriteTimeout:      writeTimeout,
		MaxHeaderBytes:    maxHeadBytes,
	}

	server.ListenAndServe()

}
