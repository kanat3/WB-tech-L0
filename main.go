package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	fmt.Println(cfg.ServerHost + ":" + cfg.ServerPort)

	router = gin.Default()

	router.LoadHTMLFiles(cfg.HTML + "index.html")

	router.Static("/images", "./images")

	router.GET("/", index)

	router.GET("/images", index)

	router.Run(cfg.ServerHost + ":" + cfg.ServerPort)

}

func index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}
