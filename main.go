package main

import "github.com/gin-gonic/gin"

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		IP := c.ClientIP()
		c.String(200, IP)
		})
	r.Run(":80") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}