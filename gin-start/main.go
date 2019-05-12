package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/a.json", func(context *gin.Context) {

	})
	router.POST("/a.json", func(context *gin.Context) {

	})
	router.PUT("/a.json", func(context *gin.Context) {

	})
	router.DELETE("/a.json", func(context *gin.Context) {

	})
	router.Run(":8887")
}
