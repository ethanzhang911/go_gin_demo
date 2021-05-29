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
	r.Run()
}
