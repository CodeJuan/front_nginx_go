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

func init(){
	Users = make(map[string]User)
}

func name(c *gin.Context){
	name := c.Param("name")
	fmt.Println(name)
	val, ok := Users[name]
	fmt.Println(ok)
	if(ok == true){
		c.JSON(http.StatusOK, gin.H{"name":name, "age":val.Age})
	}else{
		c.JSON(http.StatusNotFound, gin.H{})
	}

}

func GetUsers(c *gin.Context) {
	var tmpUsers []User
	for _, user := range Users{
		tmpUsers = append(tmpUsers, user)
	}
	c.JSON(http.StatusOK, tmpUsers)
}

func post(c *gin.Context){
	fmt.Println(c)
	fmt.Println(*c)
	var json User
	if c.BindJSON(&json) == nil {
		fmt.Println(json)
		Users[json.Name] = json
		c.JSON(http.StatusOK, gin.H{"status": "added"})
		fmt.Println(Users)
	}else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error"})
	}
}

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
	}
	s.Run("0.0.0.0:80")
}