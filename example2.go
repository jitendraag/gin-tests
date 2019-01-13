package main

import "github.com/gin-gonic/gin"

// HTTP verbs
func main() {
	gin.DisableConsoleColor()

	router := gin.Default()

	router.GET("/someGet", handleAll)
	router.POST("/somePost", handleAll)
	router.PUT("somePut", handleAll)
	router.DELETE("someDelete", handleAll)
	router.PATCH("somePatch", handleAll)
	router.HEAD("someHead", handleAll)
	router.OPTIONS("someOptions", handleAll)

	router.Run(":3000")
}

func handleAll(context *gin.Context) {
	context.JSON(200, gin.H{
		"working": "yes",
	})
}