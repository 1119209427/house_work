package api

import (
	"business-incqupt/model"
	"business-incqupt/service"
	"business-incqupt/tool"
	"fmt"
	"github.com/gin-gonic/gin"
)

func register(ctx *gin.Context){
	username:=ctx.PostForm("username")
	password:=ctx.PostForm("password")
user:=model.User{
	UserName: username,
	PassWord: password,

}
flag,err:=service.IsRepeatUsername(username)
if err!=nil{
	fmt.Println("用户名重复",err)
	tool.RespInterError(ctx)
	return
}
if flag{
	tool.RespErrorWithDate(ctx,"用户名已经存在")
	return
}
err=service.Register(user)
if err!=nil{
	fmt.Println("register err:",err)
	tool.RespInterError(ctx)
	return
}
tool.RespSuccesful(ctx)


}
func login(ctx *gin.Context){
	username:=ctx.PostForm("username")
	password:=ctx.PostForm("password")
	flag,err:=service.IsPasswordCorrect(username,password)
	if err!=nil{
		fmt.Println("",err)
		tool.RespInterError(ctx)
	}
	if !flag{
		fmt.Println("login err:",err)
		tool.RespErrorWithDate(ctx,"密码错误")
		return

	}
	ctx.SetCookie("username",username,600,"/","",false,false)
	tool.RespSuccesful(ctx)

}
func changePassword(ctx *gin.Context) {
	oldPassword := ctx.PostForm("old_password")
	newPassword := ctx.PostForm("new_password")
	iUsername, _ := ctx.Get("username")
	username := iUsername.(string)
	flag,err:=service.IsPasswordCorrect(username,oldPassword)
	if err!=nil{
		fmt.Println("判断正确密码错误：",err)
		tool.RespInterError(ctx)
		return

	}
	if !flag{
		tool.RespErrorWithDate(ctx,"旧密码错误")
		return
	}
	err=service.ChangePassword(username,newPassword)
	if err!=nil{
		fmt.Println("修改密码失败",err)
		tool.RespInterError(ctx)
		return
	}
	tool.RespSuccesful(ctx)
}