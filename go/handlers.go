package main
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"io/ioutil"
	"strings"
	"github.com/bitly/go-simplejson"
	"bytes"
)


const API_URL = "http://113.140.71.252:9091/"

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
	name := c.DefaultQuery("name", "高新3号线")
	name = strings.TrimSpace(name)

	url := API_URL + "xa_gj_mobile_provide/getBusStartEndStationByRouteName.action?routeName=" +
			name

	contents, err := get_bus_provide(url)
	if err == nil {
		c.String(http.StatusOK, string(contents))
	}
	c.String(http.StatusOK, "")
}

func get_bus_provide(url string) ([]byte, error){
	fmt.Println(url)
	response, err := http.Get(url)
	var contents []byte
	if err != nil {
		fmt.Printf("%s", err)
		return contents, err
	} else {
		defer response.Body.Close()
		contents, err = ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			return contents, err
		}
	}
	return contents, nil
}

func get_route_by_id(c *gin.Context) {
	id := c.DefaultQuery("id", "6345")

	url := API_URL + "xa_gj_mobile_provide/getStationByRouteId.action?routeId=" + id
	route_byte, err := get_bus_provide(url)
	if err != nil {
		c.String(http.StatusOK, "")
	}

	route, err := simplejson.NewFromReader(bytes.NewBuffer(route_byte))
	if err != nil {
		c.String(http.StatusOK, "")
	}
	fmt.Println(route)

	// 这里假定up是0，down是1
//	up_url := API_URL + "xa_gj_mobile_provide/getNumberPlateStationIdByRouteIdRunningType.action?routeId="
//	up_url += id + "&runningType="
//	down_url := up_url + "0"
//	up_url += "1"
//	upBus_byte, err := get_bus_provide(up_url)
//	if err != nil {
//		c.String(http.StatusOK, "")
//	}
//	downBus_byte, err := get_bus_provide(down_url)
//	if err != nil {
//		c.String(http.StatusOK, "")
//	}
//
//	upBus, err := simplejson.NewFromReader(bytes.NewBuffer(upBus_byte))
//	if err != nil {
//		c.String(http.StatusOK, "")
//	}
//	downBus, err := simplejson.NewFromReader(bytes.NewBuffer(downBus_byte))
//	if err != nil {
//		c.String(http.StatusOK, "")
//	}

	c.JSON(http.StatusOK, route)
}