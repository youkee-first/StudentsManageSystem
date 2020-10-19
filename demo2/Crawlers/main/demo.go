package main

import (
	"demo2/Crawlers/dao"
	"demo2/Crawlers/domain"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func FindAll(c *gin.Context) {
	c.HTML(http.StatusOK,"homeu.html",dao.FindAll())
}
func AddStudent(c *gin.Context) {
	sid:=c.PostForm("id")
	name:=c.PostForm("name")
	sfloor:=c.PostForm("floor")
	sroom:=c.PostForm("room")
	id,_:=strconv.Atoi(sid)
	floor,_:=strconv.Atoi(sfloor)
	room,_:=strconv.Atoi(sroom)
	dao.AddStudent(domain.Student{id,name,floor,room})
	c.HTML(200,"homeu.html",dao.FindAll())
}

func DeleteById(c *gin.Context) {
	sid:=c.PostForm("id")
	id,_:=strconv.Atoi(sid)
	dao.DeleteById(id)
	c.HTML(200,"homeu.html",dao.FindAll())
}
func homeu(c *gin.Context) {
	c.HTML(200,"homeu.html",nil)
}

func UpdateStudent(c *gin.Context) {
	sid:=c.PostForm("id")
	name:=c.PostForm("name")
	sfloor:=c.PostForm("floor")
	sroom:=c.PostForm("room")
	id,_:=strconv.Atoi(sid)
	floor,_:=strconv.Atoi(sfloor)
	room,_:=strconv.Atoi(sroom)
	dao.UpdateStudent(domain.Student{id,name,floor,room},id)
	c.HTML(200,"homeu.html",dao.FindAll())
}

func main() {
	r := gin.Default()
	//r.Static()
	r.Static("/statics","../statics")
	r.LoadHTMLFiles("./homeu.html","./QueryRoommates.html","./roomduty.html")
	r.GET("/roomduty", func(c *gin.Context) {
		c.HTML(200,"roomduty.html",nil)
	})
	r.POST("/SetDutyDate", func(c *gin.Context) {
		s := struct {
			Id int `form:"id"`
			Year int `form:"year"`
			Month int `form:"month"`
			Day int `form:"day"`
		}{}
		_ = c.ShouldBind(&s)
		fmt.Println(s)
		dao.SetDutyDate(s.Id,s.Year,s.Month,s.Day)
		c.HTML(200,"roomduty.html",nil)
	})
	r.POST("/FindDutyStudent", func(c *gin.Context) {
		s := struct {
			Year int `form:"year"`
			Month int `form:"month"`
			Day int `form:"day"`
		}{}
		_ = c.ShouldBind(&s)
		fmt.Println(s)
		c.HTML(200,"roomduty.html",dao.FindDutyStudent(s.Year,s.Month,s.Day))
	})
	r.GET("/homeu",homeu)
	r.GET("/FindAll",FindAll)
	r.POST("/AddStudent",AddStudent)
	r.POST("/DeleteById",DeleteById)
	r.POST("/UpdateStudent",UpdateStudent)
	r.GET("/Query", func(c *gin.Context) {
		c.Request.URL.Path = "/QueryRoommates"
		r.HandleContext(c)
	})
	r.GET("/QueryRoommates", func(c *gin.Context) {
		c.HTML(200,"QueryRoommates.html",nil)
	})
	r.POST("/QueryRoommatesById", func(c *gin.Context) {
		sid:=c.PostForm("id")
		id,_:=strconv.Atoi(sid)
		c.HTML(200,"QueryRoommates.html",dao.FindRoommateById(id))
	})
	r.POST("/QueryRoommatesByName", func(c *gin.Context) {
		name:=c.PostForm("name")
		c.HTML(200,"QueryRoommates.html",dao.FindRoommateByName(name))
	})
	r.POST("/QueryRoommatesByTwo", func(c *gin.Context) {
		sfloor:=c.PostForm("floor")
		sroom:=c.PostForm("room")
		floor,_:=strconv.Atoi(sfloor)
		room,_:=strconv.Atoi(sroom)
		c.HTML(200,"QueryRoommates.html",dao.FindRoommateByNum(floor,room))
	})
	r.Run(":8888") // listen and serve on 0.0.0.0:8080
}
