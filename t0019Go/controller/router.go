package controller

import (
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("view/*.html")

	r.GET("/", getTop)
	r.GET("/signup", getSignup)
	r.POST("/signup", postSignup)
	r.GET("/login", getLogin)
	r.POST("/login", postLogin)

	//自分の開発用

	// encrypt
	// パスワード生成
	r.GET("/test", gettestEnPw)
	r.POST("/test", posttestEnPw)
	//パスワード一致確認
	r.GET("/test2", gettestEnPw2)
	r.POST("/test2", posttestEnPw2)

	// ランダム値生成
	// 32byte
	r.GET("/test3", get32byte)
	// router.POST("/test3", posttestEnPw2)

	aes := r.Group("/test4_aes")
	{
		aes.GET("/", aesHome)
		aes.GET("/encrypt", toAes)
		aes.GET("/decrypt", deAes)
	}
	return r
}
