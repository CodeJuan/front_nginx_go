package main

type User struct {
	Name string `json:"name"  binding:"required"`
	Age  int32 `json:"age"  binding:"required"`
}

var Users map[string]User

type ReturnData struct {
	DownData []Station	`json:"downData"`
	UpData []Station	`json:"upData"`
}

type Station struct{
	StationID int64
	StationName string
	RunningType int
	OrderNumber int64
	CountOfBus	int
}

type DataType struct{
	DownData []Station	`json:"downData"`
	UpData []Station	`json:"upData"`
}

type RouteData struct {
	Message string
	TotalCount int
	Status string
	Data []DataType	`json:"data"`
}