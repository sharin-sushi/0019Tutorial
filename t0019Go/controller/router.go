package controller

import (
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("view/*.html")

	router.GET("/", getTop)
	router.GET("/signup", getSignup)
	router.POST("/signup", postSignup)
	router.GET("/login", getLogin)
	router.POST("/login", postLogin)

	//自分の開発用

	// パスワード生成
	router.GET("/test", gettestEnPw)
	router.POST("/test", posttestEnPw)

	//パスワード一致確認
	router.GET("/test2", gettestEnPw2)
	router.POST("/test2", posttestEnPw2)

	return router
}
