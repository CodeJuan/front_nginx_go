package main

import (
	"net/http"
	_ "log"
	"github.com/gin-gonic/gin"
)


func main(){
	s := gin.Default()

	s.StaticFile("/", "../front/index.html")
	s.StaticFile("/favicon.ico", "../front/favicon.ico")
	s.StaticFS("/assets", http.Dir("../front/assets"))

	v1 := s.Group("/api/v1")
	{
		v1.GET("/ping", pingPong)
		v1.GET("/user/:name", name)
		v1.GET("/users", GetUsers)
		v1.POST("/user", post)
		v1.GET("/routes", get_route_by_name)
	}
	s.Run("0.0.0.0:80")
}