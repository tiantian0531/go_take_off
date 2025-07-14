package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){
	r:=gin.Default()
	r.GET("/hello",Hello)
	// 设置监听端口为 8080，并启动服务器
    r.Run(":8080")
}

func Hello(c *gin.Context){

	c.String(http.StatusOK,"hello world")

}