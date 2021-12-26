package api

import (
	"business-incqupt/tool"
	"github.com/gin-gonic/gin"
)


func auth(ctx *gin.Context){
	username,err:=ctx.Cookie("username")
	if err!=nil{
		tool.RespErrorWithDate(ctx,"用户名错误")
		return
	}
	ctx.Set("username",username)
	ctx.Next()

}
