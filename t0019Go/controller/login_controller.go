package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sharin-sushi/0019testGin.git/crypto"
	"github.com/sharin-sushi/0019testGin.git/model"
)

func getSignup(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

func postSignup(c *gin.Context) {
	id := c.PostForm("user_id")
	pw := c.PostForm("password")
	// fmt.Printf("id=%v, pw=%v \n", id, pw)
	user, err := model.Signup(id, pw)
	fmt.Printf("user=%v, err=%v \n", user, err)

	if err != nil {
		c.Redirect(301, "/signup")
		return
	}
	c.HTML(http.StatusOK, "home.html", gin.H{"user": user})
}

func getLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func postLogin(c *gin.Context) {
	id := c.PostForm("user_id")
	pw := c.PostForm("password")
	fmt.Println("初期値 id=", id, "pw=", pw)
	user, err := model.Login(id, pw)
	if err != nil {
		c.Redirect(301, "/login")
		return
	}
	c.HTML(http.StatusOK, "home.html", gin.H{"user": user})
}

// ―――――――――――――――――
//別件の開発用
// test2へ接続
func gettestEnPw2(c *gin.Context) {
	c.HTML(http.StatusOK, "test2.html", nil)
}

func posttestEnPw2(c *gin.Context) {
	pw := c.PostForm("password")
	enPw := c.PostForm("enpassword")
	err := crypto.CompareHashAndPassword(enPw, pw)

	var result string
	if err != nil {
		// エラーがあればエラーを
		result = err.Error()
	} else {
		// エラーがなければ"一致"を
		result = "一致"
	}
	//一致していればtrueが返ってくる
	fmt.Println("予想pw", pw, "登録ハッシュ値", enPw, "判定", result)

	c.HTML(http.StatusOK, "test2.html", gin.H{"pw": pw, "enPw": enPw, "result": result})
}

// encriptは毎回違うハッシュ値が出るので意味が無かった
// test1へ接続
func gettestEnPw(c *gin.Context) {
	c.HTML(http.StatusOK, "test.html", nil)
}
func posttestEnPw(c *gin.Context) {
	pw := c.PostForm("password")
	fmt.Println("入力値 pw=", pw)
	enPw, _ := crypto.PasswordEncrypt(pw) //bcript値が返ってくる
	fmt.Println("変換後 pw=", enPw)

	c.HTML(http.StatusOK, "test.html", gin.H{"pw": pw, "enPw": enPw})
}
