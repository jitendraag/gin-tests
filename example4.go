package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Querystring parameters
func main() {
	router := gin.Default()

	router.GET("/welcome", func(context *gin.Context) {
		firstname := context.DefaultQuery("firstname", "Guest")
		lastname := context.Query("lastname") // convenient way to write context.Request.URL.Query().Get("lastname")

		context.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	router.Run(":8000")
}
