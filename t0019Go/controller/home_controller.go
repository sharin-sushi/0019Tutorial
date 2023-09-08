package controller

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func getTop(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", nil)
}

const test01_users = "test01_users" //Mysql側のtable名

var tmpl *template.Template

func getUserList() {
	funcMap := template.FuncMap{
		"nl2br": func(text string) template.HTML {
			return template.HTML(strings.Replace(template.HTMLEscapeString(text), "\n", "<br />", -1))
		},
	}
	tmpl, _ = template.New("article").Funcs(funcMap).ParseGlob("web/template/*")
}

//↓gormを使って省略可　selected.Nextとか
// func Index(w http.ResponseWriter, r *http.Request) {
// 	selected, err := model.Dbsql.Query("SELECT * FROM karaokelist")
// 	// fmt.Printf("selected=%v\n, err=%v\n", selected, err) //確認用
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	//変数(構造体？)selectedにkaraokelistテーブルの全てが代入された

// 	tabeleData := []Test01_users{} //懐かしのスライス
// 	for selected.Next() {
// 		kList := test01_users{}
// 		err = selected.Scan(&kList.Usere_id, &kList.password)
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		tabeleData = append(tabeleData, kList)
// 		//tableDataスライスにkList(要素)を組み込んでる…？
// 	}

// 	selected.Close()

// 	tmpl, err := template.ParseFiles("index.html")
// 	if err != nil {
// 		fmt.Printf("Failed to parse template: %s", err)
// 		log.Fatal(err)
// 	}

// 	if err := tmpl.Execute(w, tabeleData); err != nil {
// 		fmt.Printf("Failed to execute template: %s", err)
// 		log.Fatal(err)
// 	}

// }
