package main

import (
	"demo2/Crawlers/dao"
	"demo2/Crawlers/domain"
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
	r.LoadHTMLFiles("./homeu.html")
	r.GET("/homeu",homeu)
	r.GET("/FindAll",FindAll)
	r.POST("/AddStudent",AddStudent)
	r.POST("/DeleteById",DeleteById)
	r.POST("/UpdateStudent",UpdateStudent)
	r.Run(":8888") // listen and serve on 0.0.0.0:8080
}
