package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	service "github.com/lucaspolo/presunto/internal"
)

func main() {
	r := gin.Default()
	r.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	r.GET("/qcode/:q", func(c *gin.Context) {
		q := c.Param("q")
		fmt.Println(q)
		value, err := service.GetQCode(c, q)
		if err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"q":     q,
			"value": value,
		})
	})

	r.Run()
}
