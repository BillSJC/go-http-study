package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type data struct {
	Name string
	Age  int
}

var d *data

func main() {

	d = new(data)

	router := gin.Default()
	//获取信息
	router.GET("/a.json", func(c *gin.Context) {
		c.JSON(200, d)
	})
	//创建
	router.POST("/a.json", func(c *gin.Context) {
		//绑定payload中的JSON
		err := c.BindJSON(d)
		if err != nil {
			c.String(500, "%s", fmt.Sprint(err))
		} else {
			c.Status(204)
		}
	})
	router.PUT("/a.json", func(c *gin.Context) {
		err := c.BindJSON(d)
		if err != nil {
			c.String(500, "%s", fmt.Sprint(err))
		} else {
			c.Status(204)
		}
	})
	//删除
	router.DELETE("/a.json", func(c *gin.Context) {
		d = new(data)
		c.Status(204)
	})
	router.Run(":8887")
}
