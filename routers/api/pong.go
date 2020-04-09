package api

import (
	"gravitation/pkg/err_msg"
	"gravitation/pkg/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Pong(c *gin.Context)  {
	appG := app.Gin{C: c}

	appG.Response(http.StatusOK, err_msg.SUCCESS, "pong")

}