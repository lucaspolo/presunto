package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lucaspolo/presunto/docs"
	service "github.com/lucaspolo/presunto/internal"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()
	docs.SwaggerInfo.Title = "Swagger Presunto API"
	docs.SwaggerInfo.Description = "Presunto API Server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r.GET("/healthcheck", Healthcheck)

	r.GET("/qcode/:q", Qcode)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	if err := r.Run(); err != nil {
		panic(err)
	}
}

// ShowAccount 	 godoc
// @Summary      Show a healthcheck message
// @Description  get healthcheck message
// @Tags         healthcheck
// @Accept       json
// @Produce      json
//
// @Router       /healthcheck [get]
func Healthcheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

// ShowAccount 	 godoc
// @Summary      Show a Qcode description
// @Description  get qcode by code
// @Tags         qcode
// @Accept       json
// @Produce      json
// @Param        qcode         path      string  true  "QCode"
//
// @Router       /qcode/{qcode} [get]
func Qcode(c *gin.Context) {
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
}
