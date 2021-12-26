package api

import "github.com/gin-gonic/gin"

func InitEngine(){
	engine:=gin.Default()
	engine.POST("/register",register)
	engine.POST("login",login)
	userGroup:=engine.Group("/user")
	{
		userGroup.Use(auth)
		userGroup.POST("/password",changePassword)
	}
}
