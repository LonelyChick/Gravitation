package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gravitation/pkg/logging"
	"log"

	"net/http"

	"gravitation/pkg/setting"
	"gravitation/pkg/util"
	"gravitation/routers"
)

func init() {
	setting.Setup()

	logging.Setup()
	util.Setup()
}

func main() {
	gin.SetMode(setting.ServerSetting.Runmode)

	routerInit := routers.Router()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	listenPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeadBytes := 1 << 20

	server := &http.Server{
		Addr:              listenPoint,
		Handler:           routerInit,
		ReadTimeout:       readTimeout,
		WriteTimeout:      writeTimeout,
		MaxHeaderBytes:    maxHeadBytes,
	}

	log.Printf("[info] start http server listening %s", listenPoint)
	server.ListenAndServe()

}
