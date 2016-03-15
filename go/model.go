package main

type User struct {
	Name string `json:"name"  binding:"required"`
	Age  int32 `json:"age"  binding:"required"`
}

var Users []User
