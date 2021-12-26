package tool

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RespErrorWithDate(ctx *gin.Context,data interface{}){
	ctx.JSON(http.StatusOK,gin.H{
		"info":data,
	})
}
func RespInterError(ctx *gin.Context){
	ctx.JSON(http.StatusAlreadyReported,gin.H{
		"info": "服务器错误",
	})
}
func RespSuccesful(ctx *gin.Context){
	ctx.JSON(http.StatusOK,gin.H{
		"info":"成功",
	})
}

