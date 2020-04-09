package app

import (
	"github.com/gin-gonic/gin"
	"gravitation/pkg/err_msg"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int			`json:"code"`
	Msg string			`json:"msg"`
	Data interface{}	`json:"data"`
}

func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  err_msg.GetMsg(errCode),
		Data: data,
	})

}