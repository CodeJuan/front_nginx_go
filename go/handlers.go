package main
import (
	"github.com/gin-gonic/gin"
	"net/http"
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

func get_route_by_name(c *gin.Context){
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{"name": name})
}
