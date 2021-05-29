package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
       c.JSON(200,gin.H{
       	"message":"pong",
	   })
	})
	//get方法
	r.GET("/get", func(c *gin.Context) {
		c.String(200,"get")
	})
	r.POST("/post", func(c *gin.Context) {
		c.String(200,"post")
	})

	//注意这里httpMethod是大写DELETE
	r.Handle("DELETE","/delete", func(c *gin.Context) {
		c.String(200,"delete")
	})
	// 启动gin服务，默认情况下是8080端口，当然也可以自定义端口号
	r.Run("127.0.0.1:80")
}
