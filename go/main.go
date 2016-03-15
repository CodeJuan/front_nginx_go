package main

import (
	"net/http"
	_ "log"
	"github.com/gin-gonic/gin"
	"fmt"
)

func pingPong(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{"message":"pong"})
}

func name(c *gin.Context){
	name := c.Param("name")
	fmt.Println(name)
	c.JSON(http.StatusOK, gin.H{"name":name, "age":12})
}

func post(c *gin.Context){
	var json User
	if c.BindJSON(&json) == nil {
		fmt.Println(json)
		Users = append(Users, json)
		c.JSON(http.StatusOK, gin.H{"status": "added"})
		fmt.Println(Users)
	}else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error"})
	}
}

func main(){
	s := gin.Default()

	v1 := s.Group("/api/v1")
	{
		v1.GET("/ping", pingPong)
		v1.GET("/user/:name", name)
		v1.POST("/user", post)
	}
	s.Run("0.0.0.0:80")
}