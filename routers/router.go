package routers

import (
	"gravitation/routers/api"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine  {
	//r := gin.New()
	r := gin.Default()
	r.GET("/ping", api.Pong)

	return r
}