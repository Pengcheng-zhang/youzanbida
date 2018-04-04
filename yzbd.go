package main

import (
	"fmt"
	"time"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"biz"
	"route"
	//"yztest"
)

func main()  {
	now,err := time.Parse("2006-01-02 15:04:05 +0000 UTC", "2018-01-09 11:02:51 +0000 UTC" )
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	biz.DbInit();
	defer biz.GetDbInstance().Close()
	//var comBiz biz.CommomBiz
	//result := comBiz.GetCategory(2)
	//var length int = len(result)
	//cateIds := make([]int, length)
	//yztest.Run()
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Directory:"templates",
		Layout: "layout",
		Extensions:[]string{".tmpl",".html"},
		Charset:"UTF-8",
		IndentJSON: true,
	}))
	store := sessions.NewCookieStore([]byte("secret123"))
	m.Use(sessions.Sessions("my_session", store))
	route.Run(m)
}
